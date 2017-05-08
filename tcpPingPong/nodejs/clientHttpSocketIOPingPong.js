// How to use:
//   node clientHttpSocketIOPingPong.js <ip:port>

var httpsvr;
if(process.argv[2]){
  httpsvr = 'http://'+process.argv[2];
} else {
  httpsvr = 'http://127.0.0.1:8010/';
}

var nproc;
if(process.argv[3]){
  nproc = parseInt(process.argv[3]);
} else {
  nproc = 1;
}

var io = require('socket.io-client');
var stdin = process.openStdin();

var spawnio = function (delay){
  var socket = io.connect(httpsvr, {reconnect: true});
  var q = [];
  q.socket = socket;
  //delay = Math.ceil(Math.random()*10000);
  socket.on('connect', function (s) {
    console.log('Connected! '+q.socket.io.engine.id);
    var emit = function(){
      msg = 'id: '+q.socket.id+' delay: '+delay.toString()+' msg: kuda dikirim tanggal '+Date();
      //q.socket.emit('ch1', msg);
      //console.log('send: '+msg);
    }
    setInterval(emit,delay);
  });
};

for(var i= 0; i<nproc; i++) {
  delay = Math.ceil(Math.random()*10000);
  spawnio(delay);
}

