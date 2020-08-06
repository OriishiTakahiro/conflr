import React from "react";
import { Facet, Article } from "./Models";

// styling
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Link from "@material-ui/core/Link";
import { Button } from "@material-ui/core";
import { createStyles, Theme, makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";

type ResultProps = {
  articles: Article[];
  facets: Facet[];
  filter: string[];
  setFilter: React.Dispatch<React.SetStateAction<string[]>>;
};

const newFacetStyle = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: "100%",
      maxWidth: 360,
      backgroundColor: theme.palette.background.paper,
    },
  })
);

const newArticleStyle = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: "100%",
      backgroundColor: theme.palette.background.paper,
    },
  })
);

export function Results(props: ResultProps) {
  const facetStyle = newFacetStyle();
  const articleStyle = newArticleStyle();
  return (
    <div className="Result">
      <Grid
        container
        direction="row"
        justify="flex-start"
        alignItems="flex-start"
      >
        <Grid item xs={2}>
          <List className={facetStyle.root}>
            {props.facets.map((e, i) => (
              <ListItem
                key={i}
                alignItems="flex-start"
                onClick={(event) => props.setFilter(props.filter.concat(e.val))}
              >
                <Button color="primary">
                  {e.val} ({e.count})
                </Button>
              </ListItem>
            ))}
          </List>
        </Grid>
        <Grid item xs={10}>
          <List className={articleStyle.root}>
            {props.articles.map((e) => (
              <ListItem key={e.id} alignItems="flex-start">
                <ListItemText
                  primary={
                    <Link href={e.url}>{`${e.title}: ${e.displayName}`}</Link>
                  }
                  secondary={
                    typeof e.view === "undefined"
                      ? "no contents"
                      : e.view.substring(0, 400) + "..."
                  }
                ></ListItemText>
              </ListItem>
            ))}
          </List>
        </Grid>
      </Grid>
    </div>
  );
}
