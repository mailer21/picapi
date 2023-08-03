# WebP Converter API

This is a simple Go API that accepts images in .jpeg, .jpg, or .png formats and converts them to .webp format, with the ability to control the compression ratio.

## How to Run

1. Make sure you have Go installed (version 1.15 and above).

2. Clone the repository:
   ``` bash
   git clone https://github.com/mailer21/picapi.git
   ```
3. Go to the project directory:
    ``` bash
    cd picapi
    ```
4. Use environment variables
    ``` bash
   PORT=8080 # Port on which the server will run (default is 8080).
   COMPRESS_RATIO=80 # Compression ratio for the converted images (default is 80).
   ```
5. Run the server:
   ``` bash
    go run ./main.go
    ```
6. The server will be running at `http://localhost:8080`.

## How to Use with UI

1. Open a web browser and go to `http://localhost:8080/ui`.

2. Choose a file in .jpeg, .jpg, or .png format that you want to convert to .webp.

3. Click on the "Convert to WebP" button.

4. After processing the request, a link to download the converted .webp file with the original filename will appear below the uploaded image.

## How to Use API
URL: /convert 
Method: POST
Request Body: The image file to be converted in a multipart/form-data format, with the key name as image.
Example Request:
``` bash
curl -X POST -F "image=@example.jpg" http://localhost:8080/convert
```
Example Response:
``` bash
{
  "fileName": "example.webp",
  "dataURL": "data:image/webp;base64,<data>
}
```

## Examples

### Example 1

Filename: `example.jpg`

File Size: 1000 KB

Compression Ratio: 80%

Result: `example.webp`

### Example 2

Filename: `image.png`

File Size: 500 KB

Compression Ratio: 70%

Result: `image.webp`

## Notes

* The application only supports images in .jpeg, .jpg, and .png formats.
* The compression ratio for the .webp file is determined by the `COMPRESS_RATIO` environment variable in the `.env` file and can range from 1 to 100. Higher values result in less compression and larger file sizes.

I hope this README.md helps you to use and set up the application for converting images to .webp format. If you have any questions or suggestions, please feel free to ask!
