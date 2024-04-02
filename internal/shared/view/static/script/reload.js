(function(){
  const conn = new WebSocket("ws://" + document.location.host + "/livereload");  
  conn.onclose = function (evt) {  
    console.log("Connection closed.. Reloading.")  
    setTimeout(function () {  
      location.reload();
    }, 2000);
  };

  conn.onopen = function (evt) {
    console.log("Connection open.");
  }
})()