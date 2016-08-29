# City
A simple in memory feed aggregator service


## Running local

You need a configuration file on yaml format with your feed title and a
list of feed urls like:

```yaml
title: "Test Config Title"
feeds:
    - "https://raelmax.github.io/rss.xml"
```

Then, download the binary on [releases page](https://github.com/raelmax/city/releases) and run:
```
./city -config=path/to/your/config.yaml -port=8000 -timeout=5
```

Access [http://127.0.0.1:8000](http://127.0.0.1:8000) from your browser.


## Deploy to heroku
Just fork this repository and click on this button:

[deploy to heroku button]

## Development
Get the _city_ source code:
```
go get github.com/raelmax/city
```

Run tests:
```
go test
```

## Future
- [ ] Deploy to heroku
- [ ] Improve test coverage
- [ ] Pagination
- [ ] Posts expiration
- [ ] Database support(?)