**To do:**

- [x] Database
    - [x] Create User schema
    - [x] Create Message schema
    - [x] Create Room Schema
    - [x] Room user junction table

- [ ] API
    - [x] Users (get all)
        - [x] Login (JWT implementation)
        - [x] Add (register)
            - [ ] Confirmation
    - [x] Message (get by room id, user id)
        - [x] Send message by user id and room
        - [ ] Delete message
    - [x] Room (get by id, all)
        - Create room

- [x] Authentication (JWT)

- [ ] Websockets to check online state

- [ ] Web interface
    - [x] Login, JWT storage
        - [ ] Register
        - [ ] Room access permissions
    - [x] Home page (just render rooms to table)
    - [ ] Room page
        - [ ] Message refresh on send
        - [ ] Limit amount of messages to render
            - [ ] Control messages array using state

- [ ] Tests

**Business functions notes:**

- User login, lookup, register
- Chat room create, delete, join
    - Search for users
- Message get, send, delete for room.

**Log:**

- Learning stuff. 
- Repository pattern for multiple objects (allows switching databases & testing), basic http server routing.  
- Test strategy pattern, research api design, app infrastructure, ? how to join tables in repository patterns.  
- 12/12: Implement auth.
- 16/12: Test nextjs, react. Now using vite + react, testing jwt storage on browser, request cross origin using cors header.
- 17/12: Practice useEffect.
- ~20/12: still improving ui, adding functions, restructuring code.
- 21/12: failed to implement message update, need to look into **websockets** for real time update.
- ~23/12: done half websocket implementation (go server), need authentication over websockets which seems troublesome. Might switch to pubsub later. Current solution is smuggling auth tokens over protocol.

**Resource:**

- https://dev.to/santoshanand/implementing-the-repository-pattern-in-go-with-both-in-memory-and-mysql-repositories-581j
- https://github.com/jorzel/go-repository-pattern
- https://www.alexedwards.net/blog/organising-database-access (\*)
- https://www.alexedwards.net/blog/making-and-using-middleware (\*)
- https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor (\*)


- https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/
- https://github.com/golang-jwt/jwt
- https://permify.co/post/jwt-authentication-go/ (\*)


- https://stackoverflow.com/questions/35553500/xmlhttprequest-cannot-load-xxx-no-access-control-allow-origin-header (\*)
- https://www.descope.com/blog/post/developer-guide-jwt-storage


- https://stackoverflow.com/a/77060459/31690738
- https://jsmanifest.com/the-publish-subscribe-pattern-in-javascript/ (future reading)
