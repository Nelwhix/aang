# AANG
A CLI tool to prepare my nuxt app for staging. It does the following:
- git add
- git commit
- copy the commit id, overwrite the .env and up the version number and change the commit id
- npm run generate
- git push to the dev repo
- git push to the staging repo

## USAGE
- clone the repo
- https://go.dev/doc/tutorial/compile-install

To see all the command line flags and options
```bash
    aang -h
```