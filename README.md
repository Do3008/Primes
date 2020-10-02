# # Documentation

## Authentication

### 1. Set session `https://primes-mulapese.herokuapp.com/set_session` (GET)

### 2. Get session `https://primes-mulapese.herokuapp.com/get_session` (GET)
**Response**
- 500 Server error
- 200 OK - {status, session}
### 3. Get session `https://primes-mulapese.herokuapp.com/destroy_session` (GET)
**Response**
- 500 Server error
- 200 OK - {status, session}

## Primes API
### 1. Get sum of primes below N `https://primes-mulapese.herokuapp.com/primes/sum/:num` (GET)
**Request**
- param num is a number greater than or equal to 2
- session (have granted before)

**Response**
- 500 Server error
- 401 Unauthentication
- 400 Bad request - wrong param (number greater than or equal to 2)
- 200 OK - response object contain sum of all primes below n {sum}

# # User manual
## Sum all primes below N
### Step 1: Set session by go to `https://primes-mulapese.herokuapp.com/set_session` (GET)
### Step 2:  Get sum of primes below N by go to `https://primes-mulapese.herokuapp.com/primes/sum/:num` (GET) and replace :num with N number you want

# # Installation guide
## Nodejs (All code in vinid folder)
- Clone project from this github https://github.com/Do3008/Primes.git
- Open cmd and go to vinid folder
- Install nodejs by follow direction in https://nodejs.org/en
- Install all dependencies with ```npm install``` in command line
- Run program with ```npm start```
- To test program run ```npm test```
## Golang (All code in vinid-golang folder)
- Open cmd and go to vinid-golang folder
- Install golang by follow direction in https://golang.org/dl/
- Run program with ```go run main.go```

# # Deployment guide
- Open cmd and go to vinid folder and follow these step
- Initialize a local Git repository ```git init``` ```git add .```  ```git commit -m``` ```"first deploy"```
- Create heroku remote ```heroku create primes-mulapese``` ```git remote -v``` ```heroku git:remote -a primes-mulapese```
- Deploying code ```git push heroku master```
- go to https://primes-mulapese.herokuapp.com/ to view the website
