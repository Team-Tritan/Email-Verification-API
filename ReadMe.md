# Functional, but not completed. :)

# Email Verification Service

A very **very** basic email verif microservice made for a friend's project. The motivation behind making this was to better my own skills with Go. I've always handled email verification either in the frontend or the registration backend workflow so it's an interesting take on handling it differently, independantly. You can also use this service with any frontend, it simply will return if the user has verified their email or not upon checking.

### API Routes:

#### Base URL

- GET / - What do you expect? It's an api landing page

#### Send Verification Request

- GET /api/send/[email address]?token=[api key]
  - Store `id` that is returned to you, correlate with user in your own db. This is their verification token and also how you look up the status of their verification.

#### Get Verification Status

- GET /api/check/[id]?token=[api key]
  - Use the stored `id` provided in res for the the `/api/send/email` route that you correlated with your user to check for status. Check on a loop until `verified: true` is returned to you.

#### Verify Email Route

- GET /api/verify/[id]
  - This is the route that actually verifies the user when clicking the link in the email sent to them.

#### Get API Status

- GET /api/status
  - Use uptime monitor to check this route specifically.

### UI Images:

![email](https://i.imgur.com/KwGKnad.png)
