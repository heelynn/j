# j 

[中文wiki](https://github.com/heelynn/j/blob/main/README_ZH.md)

Note: The `master` branch may be under development and is not a stable version. Please download the stable version of the source code via tags, or download the compiled binary executable files from the [release](https://github.com/heelynn/j/releases).

`j` is a command-line tool for Linux, macOS, and Windows that provides a convenient way to manage and switch between multiple Java environments.

`j` can easily switch Java versions.

## Manual Installation
- Windows
```shell
# Using PowerShell  
mkdir $HOME\.j\bin, $HOME\.j\version  
# Choose the Windows version to download the installer from: https://github.com/heelynn/j/releases/latest  
# Extract the j.exe file to the $HOME\.j\bin directory
````
Configure environment variables. It is recommended to use the system interface to set environment variables for permanent effect.

```shell
$env:JAVA_HOME="$HOME\.j\java"  
$env:PATH="$HOME\.j\bin;$JAVA_HOME\bin;$env:PATH;"
```
- macOS/Linux
```shell
# Using shell  
$ mkdir -p $HOME/.j/bin $HOME/.j/version  
# Choose the Linux/macOS version to download the installer from: https://github.com/heelynn/j/releases/latest  
  
# Extract the j file to the $HOME/.j/bin directory  
# Add executable permissions to the j file  
$ chmod +x $HOME/.j/bin/j  
  
# Configure environment variables (for bash, zsh)  
$ export JAVA_HOME="$HOME/.j/java"  
$ export PATH="$HOME/.j/bin:$JAVA_HOME/bin:$PATH"
```
## Usage

### Download and Install Java
Extract the downloaded Java installer (e.g., Redhat, OpenJDK, OracleJDK) to the $HOME/.j/version directory.
It is recommended to rename the folders to simple names such as Java8, Java11, Java17 for easy management.

```shell
# List installed Java versions  
$ ls $HOME/.j/version  
Java8  
Java11  
Java17
```
## Check Current Java Version
`Use the j command` to check the current Java version, which corresponds to the Java versions in the $HOME/.j/version directory.

```shell
$ j ls  
java8  
java11
```
## Switch Java Version
Use the `j` command to switch the Java version. Use j ls to view the currently installed Java versions.
```shell
$ j use java8
```
Verify Java Version
```shell
$ java -version
```
