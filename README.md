### Permie

An uselessly simple in-memory role-permission library for Golang applications. 
This was created to suit applications that do not have complex permission requirements.

#### Examples
```go
permie.Define("admin", permie.Permissions{"users.create": true})
```
```go
permie.Allow("admin", "users.ban")
permie.Disallow("visitor", "posts.read")
```
```go
// imagine we are requesting user from database with a list of roles
user := ... 
if permie.CanAtLeast(user.roles, "posts.read") {
	// ... do something here
}
```
```go
user := ...
// This allows checking whether the user can read and delete posts.
if permie.Require(user.roles, "posts.read", "posts.delete") {
	// do something here
}
```
```go
if permie.Can("admin", "users.ban") {
	// do something here
}
```

#### Installation
```go
go get github.com/ShindouMihou/Permie
```