var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
const session = require('express-session');


var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');

var app = express();

app.use(session({
  resave: true, 
  saveUninitialized: true, 
  secret: 'somesecret', 
  cookie: { maxAge: 1000 * 60 * 60 * 24 }}));

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);
 
//set session
app.get('/set_session', (req, res) => {
  //set a object to session
  req.session.User = {
      website: 'anonystick.com',
      type: 'blog javascript',
      like: '4550'
  }

  return res.status(200).json({status: 'success'})
})

//set session
app.get('/get_session', (req, res) => {
  //check session
  if(req.session.User){
      return res.status(200).json({status: 'success', session: req.session.User})
  }
  return res.status(200).json({status: 'error', session: 'No session'})
})

//destroy session
app.get('/destroy_session', (req, res) => {
  //destroy session
  req.session.destroy(function(err) {
      return res.status(200).json({status: 'success', session: 'cannot access session here'})
  })
})


// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
