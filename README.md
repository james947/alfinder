# alfinder

`alfinder` is a fast and efficient command-line file search tool written in Go. It allows users to perform global file searches starting from the root directory, with options to filter by file type and search within specific directories. `alfinder` is inspired by the albatross, known for its wide-ranging travels and efficiency.

## Features

- Global file search starting from the root directory `/`.
- Option to filter search results by file type (`pdf`, `img`, `txt`).
- Prioritizes user directories over system directories for faster and more relevant results.

## Installation

- Download the package then ``go build``

## Options:

   #### Global Search for All Files Containing "report"
    sudo alfinder "report"
    

  ### Search for PDF Files Containing "invoice" in a Specific Directory:
    sudo alfinder /home/user/documents "invoice" -pdf
      


  ### Search for Image Files Containing "vacation" Globally:
    sudo alfinder "vacation" -img
    


## Contributing
- Contributions are welcome! Please fork this repository and submit a pull request with your changes.

