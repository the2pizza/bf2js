# bf2js

bf2js is a transpiler from brainf*ck to js (nodejs)

You could find examples of BF programs in ```tests/```

## Usage

By default build, trnaspiles and runs js code

```make file=<filename>```

Example: ```make file=tests/hello.b```

### Build

Builds source golang code to binary

```make build ```


### Run compiled JS

Runs transpiled js program 

```make run-node file=build/<file>.js```

### Transpile 

Transpiles bf program to js and puts file to ```build/``` 

```make transpile <file> ```


#ToDo

- Support vanilla JS for run in browser
- Add nicely handling arguments, options 
- Add more validators for BF program (at least brackets validator)

