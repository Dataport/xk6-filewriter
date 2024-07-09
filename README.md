# k6filewriter

[![Go Report Card](https://goreportcard.com/badge/github.com/Dataport/xk6-filewriter)](https://goreportcard.com/report/github.com/Dataport/xk6-filewriter)

xk6 Module for creating and writing/logging strings into local files during a loadtest with [k6](https://k6.io/). If the directories and subdirectories **do not exist**, 
the writeString() function will create them and the file with permission bits 0750. However, if they exist, the writeString() method will act like the appendString() 
method. The appendString() function will just add the respective string into the end of the file.

## Installation 

To use the module, you need to build k6 with this module. To archieve this you need [xk6](https://github.com/grafana/xk6). Otherwise in the releases section of this repository on Github, you can download a precompiled k6.exe for windows, If you do not want to use or install xk6. If you have installed xk6 it successfully, 
you can build your k6 with the module installed with the command below: 

```bash
xk6 build --with github.com/Dataport/xk6-filewriter
```

## Usage in code

To use the module in your k6 Javascript code, import the module with the following line and call the writeString() and appendString() method

```Javascript
import filewriter from "k6/x/filewriter";

//create a new file and write a string
filewriter.writeString("path", "filename", "myAddedString")

//append to an existing file
filewriter.appendString("path", "filename", "myAddedString")

// Create a new file. Existing file will be deleted
// On Windows write protection will not prevent deletion
filewriter.createFile("path", "filename")
```

## Adding new code

If you want to contribute to this projekt, please make sure that...

- [ ] that you do not use Libraries with licenses not fitting to our Apache License 2.0 License
- [ ] update the sbom accordingly 
- [ ] do not have any sensible data in the code (things like the IT-Infrastructure or personal data)
- [ ] create a merge request which gets **approved** by atleast one maintainer

*License: Apache License 2.0*

