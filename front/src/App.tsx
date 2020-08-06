import React, { useState, useEffect } from "react";
import logo from "./logo.svg";
import "./App.css";

// styling
import { Button, Input } from "@material-ui/core";

import { Search } from "./Search";
import { Results } from "./Results";

function App() {
  const [articles, setArticles] = useState([]);
  const [facets, setFacets] = useState([]);
  const [filter, setFilter] = useState(new Array<string>());
  const [query, setQuery] = useState("*");

  useEffect(() => {
    Search(query, filter, setArticles, setFacets);
  }, [filter]);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div>
          <Input
            className="QueryForm"
            color="primary"
            type="search"
            defaultValue={query}
            onChange={(e) => setQuery(e.target.value)}
          ></Input>
        </div>
        <div>
          <Button
            color="primary"
            variant="contained"
            className="Search"
            onClick={() => {
              Search(query, filter, setArticles, setFacets);
            }}
          >
            Search
          </Button>
        </div>
        <div>
          <Results
            articles={articles}
            facets={facets}
            filter={filter}
            setFilter={setFilter}
          ></Results>
        </div>
      </header>
    </div>
  );
}

export default App;
