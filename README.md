# graphql-go

GraphQL API implemented in Go and backed by MySQL

## How to run it

To run this project you need to:
- install golang, see [this guide](https://golang.org/doc/install)

## Commands

This application exposes a graphql endpoints `/graphql/v1` which accepts both mutations and queries.
The following are examples of curl calls to this endpoint:

Example:
```bash
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {createUser(email:"1@x.co"){id, email}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {createUser(email:"2@y.co"){id, email}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {follow(follower:1, followee:2)}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {unfollow(follower:1, followee:2)}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:2){followers{id, email}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){followers{id, email}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:2){follower(id:1){ email}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){followees{email}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){followee(id:2){email}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {createPost(user:1,title:"p1",body:"b1"){id}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {createComment(user:1,post:1,title:"t1",body:"b1"){id}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {removeComment(id:1)}'
curl -XPOST http://vm:8080/graphql/v1 -d 'mutation Root {removePost(id:1)}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){post(id:2){title,body}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){posts{id,title,body}}}'
curl -XPOST http://vm:8080/graphql/v1 -d 'query Root {user(id:1){post(id:2){user{id,email}}}}'

```

Other way, you can just access http://localhost:8080/ to use a playground... 

Powered with :heart: by Dropadev http://dropadev.com/
