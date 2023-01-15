# _WIP, not finished._

# Email Verification Service

A very **very** basic email verif microservice made for a friend's project. The motivation behind making this was to better my own skills with Go. I've always handled email verification either in the frontend or the registration backend workflow so it's an interesting take on handling it differently, independantly. You can also use this service with any frontend, it simply will return if the user has verified their email or not upon checking.

### API Routes:

#### Base URL

- GET / - What do you expect? It's an api landing page

#### Send Verification Request

- GET /api/send/[email address]?token=[api key]

#### Get Verification Status

- GET /api/check/[verif token]?token=[api key]

#### Verify Email Route (sent as link in email)

- GET /api/verify/[verif token]

#### Get API Status

- GET /api/status

### UI Images:

![email](https://i.imgur.com/KwGKnad.png)
