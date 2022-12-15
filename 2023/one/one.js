

function calculateFile(){
  var file = document.getElementById('file').files[0];
  var reader = new FileReader();
  reader.onload = function(progressEvent){    
    var lines = this.result.split(/\r\n|\n/);
    for(var line = 0; line < lines.length-1; line++){
      var output = document.getElementById('output');
      output.innerHTML = output.innerHTML + "\n" + lines[line];
    }
  };
  reader.readAsText(file);
};