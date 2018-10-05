# rtf

A wrapper over the amazing package that integrates go with LibreOffice 5.2

```https://github.com/dveselov/go-libreofficekit```

### Dependencies

LibreOffice 5.2 (Exact version)

### How To?

1. Install the debian package from the LibreOffice archive
https://downloadarchive.documentfoundation.org/libreoffice/old/5.0.2.2/deb/x86_64/

2. Extract the tar.gz file and install all the .deb files

3. Get the awesome go-libreofficekit 

Latest version of LibreOffice (5.2) is required
``` 
$ sudo add-apt-repository ppa:libreoffice/ppa 
$ sudo apt-get update
$ sudo apt-get install libreoffice libreofficekit-dev
$ go get github.com/docsbox/go-libreofficekit
```

4. Get this cli wrapper

``` go get github.com/minhajuddinkhan/rtf/cmd/libreconv ```

5. Convert via CLI

``` libreconv convert pathtoyourinputfile.html pathtoyouroutputfile.rtf```
