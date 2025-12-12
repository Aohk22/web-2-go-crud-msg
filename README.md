**To do:**

- [x] Database
    - [x] Create User schema
    - [x] Create Message schema
    - [x] Create Room Schema
    - [x] Room user junction table

- [ ] Functions
    - [ ] Users (get all)
        - [ ] Login (JWT implementation)
        - [x] Add (register)
            - [ ] Confirmation
    - [x] Message (get by room id, user id)
    - [x] Room (get by id, all)

- [x] Authentication (JWT)

- [ ] Websockets

- [ ] Web interface

- [ ] Tests

**Business functions notes:**

- User login, lookup, register
- Chat room create, delete, join
    - Search for users
- Message get, send, delete for room.

**Log:**

Learning stuff. 
Repository pattern for multiple objects (allows switching databases & testing), basic http server routing.  
Test strategy pattern, research api design, app infrastructure, ? how to join tables in repository patterns.  
12/12: Implement auth.

**Resource:**

- https://dev.to/santoshanand/implementing-the-repository-pattern-in-go-with-both-in-memory-and-mysql-repositories-581j
- https://github.com/jorzel/go-repository-pattern
- https://www.alexedwards.net/blog/organising-database-access (\*)
- https://www.alexedwards.net/blog/making-and-using-middleware (\*)
- https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor (\*)

- https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
- https://github.com/golang-jwt/jwt
