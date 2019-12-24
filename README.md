# go-gae-starter

A simple extensible Go API service project to run on [Google App Engine][gae] with Postgres running in [Cloud SQL][cloud-sql].

## Stack
- [Gin][gin] for HTTP routing
- [Dep][dep] for package management
- [pq][pq] for Postgres database driver
- [gcloud][gcloud] cli for deployment

[gae]: https://cloud.google.com/appengine/
[cloud-sql]: https://cloud.google.com/sql/
[gin]: https://github.com/gin-gonic/gin
[dep]: https://github.com/golang/dep
[pq]: https://github.com/lib/pq
[gcloud]: https://cloud.google.com/sdk/docs/

## Setup
Once you have cloned

- change the project path in the code. i.e. `github.com/jochasinga/boo/*` should be changed to your respective path
- run `dep ensure` in the same directory as `Gopkg.toml` to install all dependencies
- set up Postgres Cloud SQL properly and change values in `model/db.go` accordingly.
- run `gcloud app deploy`