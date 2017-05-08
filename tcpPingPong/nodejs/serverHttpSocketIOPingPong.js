// How to use:
//   node serverHttpSocketIOPingPong.js <port>

var q = [];
var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);
var path = require('path')
var dir = path.resolve()

q.con = false;
if(process.argv[2]){
  q.httpport = process.argv[2];
} else {
  q.httpport = 3000;
}

app.get('/', function(req, res){
  q.req=req;
  q.res=res;
  res.sendFile(dir+'/index.html');
});

io.on('connection', function(socket){
  q.socket = socket;
  q.con = true;
  console.log('p:'+q.httpport.toString()+' a user connected: '+socket.id);
  socket.on('disconnect', function(){
    console.log('p:'+q.httpport.toString()+' user disconnected: '+socket.id);
  });
  socket.on('ch1', function(msg){
    //q.msg = msg;
    //console.log('rcv: ' + msg);
  });
});

var main = function(){
  console.log('listening on *:'+q.httpport.toString());
}

q.http = http.listen(q.httpport, main);

setInterval(function(){
  if(q.con){
    console.log('p:'+q.httpport.toString()+' conn:'+q.socket.server.engine.clientsCount.toString());
  };
},3000);

