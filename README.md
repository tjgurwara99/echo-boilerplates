# Echo Boilerplates for my Golang based web application projects

I don't expect this to be useful to anyone else other than me but I am
making this because I have repeated this at least 4 times over the last 4
months. So hopefully this would save me some time, and if it helps anyone
else in the process, that would be a bonus!

## To get started

Install the `gonew` command

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

To create a new project using the `basic-mvt` project template, type the
following:

```sh
gonew github.com/tjgurwara99/echo-boilerplates/basic-mvt example.com/webapp
cd ./webapp
```

And then you modify the contents of the `webapp` directory for customising
it to your needs. Enjoy!
