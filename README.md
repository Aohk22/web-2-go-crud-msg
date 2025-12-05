**To do:**

- [x] Database
    - [x] Create User schema
    - [x] Create Message schema
    - [x] Create Room Schema
    - [x] Room user junction table

- [ ] Functions
    - [ ] Users
        - [ ] Login
        - [x] Add (register)
            - [ ] Confirmation
        - [x] Get all
    - [ ] Message
    - [ ] Room

- [ ] Authentication

**Business functions notes:**

- User login, lookup, register
- Chat room create, delete, join
    - Search for users
- Message get, send, delete for room.

**Log:**

Day 1-2: Learning stuff.  
Day 3: Repository pattern for multiple objects (allows switching databases & testing), basic http server routing.  
Day 4: Test strategy pattern, research api design.  

**Resource:**

- https://dev.to/santoshanand/implementing-the-repository-pattern-in-go-with-both-in-memory-and-mysql-repositories-581j
- https://github.com/jorzel/go-repository-pattern
- https://www.alexedwards.net/blog/organising-database-access (\*)
- https://www.alexedwards.net/blog/making-and-using-middleware (\*)
- https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#the-newserver-constructor (\*)
