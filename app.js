//RUNNING WITH EXTERNAL NODEJS APP


const { exec } = require('child_process');
exec('go run cmd/libreconv/main.go convert', (err, stdout, stderr) => {
 if (err) {
   console.log('errr', err)
    // node couldn't execute the command
   return;
 }

 // the entire stdout and stderr (buffered)
 console.log(`stdout: ${stdout}`);
 console.log(`stderr: ${stderr}`);
});