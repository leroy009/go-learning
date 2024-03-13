# Extending go lang functionality using 3rd party libraries

Go's standard library has so many features. you can use it to build almost everything you can think of.
But there are times when you need to perform complex and/or specialized tasks or functionality, this is where third party libraries comes in and offer you pre-built solutions for your needs.

In this post I will provide examples on using a go land thrid pary library.

## Finding the Library

Go ecosystem has a vast collection of librarues on different topics, such as database interaction, file interaction, web frameworks, machine learning, and airtificial intelligence.
You can visit the official Go Package repository [Go Packages](https://pkg.go.dev/).

Follow the below steps to use third party libraries:

1. Run the `go get` command in your terminal to download and install third pary libraries.

```bash
go get <package-url>
```

2. Import and use the third pary library.
   Once the libray has been downloaded and installed. Import the library in the file you want to use, using the following code:

    ```go
    import "<package-url>"
    ```

    After importing it. you can start using it in the file you have imported it into.

## Conclusion

Third party libraries are a valuable asset which you can use on your Go developer's toolkit. By following these guideline, you can effectively lerverage the power of these libraries to build complex, efficient and feature-rich Go applicaions. If you cannot find the package you are looking for and you can build it. Consider contributing you library to the Go community. Happy coding.
