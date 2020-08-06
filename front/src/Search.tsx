import { Facet, Article } from "./Models";
import React from "react";

const solrEndpoint = "http://localhost:8983/solr/confl/query?sow=true";

const requestTemplate = {
  query: "*",
  filter: new Array<string>(),
  offset: 0,
  limit: 10,
  fields: [
    "id",
    "title",
    "space_name",
    "createdBy_username",
    "createdBy_displayName",
    "view",
    "URL",
  ],
  facet: {
    sort: "count",
    categories: {
      type: "terms",
      field: "labels",
      limit: 10,
    },
  },
};

function filterToCondition(filter: string[]) {
  if (filter.length === 0) {
    return "";
  }
  return " AND " + filter.map((e) => `labels:${e}`).join(" AND ");
}

function weight(idx: number) {
  return 50 * Math.exp(-1 * idx);
}

function textToCondition(query: string) {
  const terms = query.split(" ");
  const title = terms.map((e, i) => `title:${e}^${weight(i)}`).join(" OR ");
  const view = terms.map((e, i) => `view:${e}^${weight(i)}`).join(" AND ");
  const space = terms.map((e, i) => `space:${e}^${weight(i)}`).join(" OR ");
  return `((${title}) OR (${view}) OR (${space}))`;
}

export function Search(
  query: string,
  filter: string[],
  setArticles: React.Dispatch<React.SetStateAction<never[]>>,
  setFacets: React.Dispatch<React.SetStateAction<never[]>>
) {
  const request = requestTemplate;
  request.query = textToCondition(query) + filterToCondition(filter);
  console.log(request.query);
  (async () => {
    await fetch(solrEndpoint, {
      method: "post",
      mode: "cors",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      cache: "no-cache",
      body: JSON.stringify(request),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        let newArticles = data.response.docs.map((e: any) => new Article(e));
        setArticles(newArticles);

        let newFacets =
          data.facets.count === 0
            ? []
            : data.facets.categories.buckets.map((e: any) => new Facet(e));
        setFacets(newFacets);
      })
      .catch(function (error) {
        console.log(error);
      });
  })();
}
