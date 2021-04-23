let readline = require('readline');
let rl = readline.createInterface(process.stdin, process.stdout);
const getChar = () => {
    return new Promise(resolve => {
        rl.on('line', (s) => {
            resolve(s);
            rl.close();
        });
    });
}

let d=new Array(30);
let i=0; let output="";
for(a=0; a<29; a++){d[a]=0}


getChar().then(char => {
    console.log(char.charCodeAt(0));
    d[0] = char.charCodeAt(0);
    console.log(d)
});

rl.on('close', function() {console.log("Closed" + d)})
while(d[i]){d[i]++;i++;i++;i++;i++;i++;i++;i++;i++;i++;i++;d[i]++;i--;i--;i--;i--;i--;i--;i--;i--;i--;i--;}
while(d[i]){i++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++;d[i]++}


console.log(d)