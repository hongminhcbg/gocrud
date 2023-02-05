# gocrud
Generate CRUD code for golang

I. Summary

  ![Alt text](./docs/gocrud.drawio.png?raw=true "Summary")

  - input is yaml file 

  - generate file magration, modeling, store

II. Install 

    Install from source code

    $ git clone https://github.com/hongminhcbg/gocrud.git

    $ cd gocrud && go mod tidy

    $ make build 

    $ make install 

III. Usage

    $ gocurd init // generate example collections.yaml, after that u update the file depend on your requirement

    $ gocrud generate // generate go code from input

### IV. TODO

- [ ] Support datetime type 
- [ ] Support Email type
- [ ] Support JSON type
- [ ] Intergrate https://github.com/go-playground/validator
