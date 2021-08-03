# Go HTTP Server

Simple HTTP server.

## Instructions

1. Run the server:

   ```sh
   go run main.go
   ```

### `/health`

1. Open <http://localhost:9000/health> in a browser or use `curl`:

   ```sh
   $ curl http://localhost:9000/health
   ok
   ```

### `/mp3`

1. Add the music file to `files`:

   ```sh
   $ mkdir files && curl --location https://freemusicarchive.org/track/Marsel_Minga_-_01_-_Demonstration/download --output files/Marsel_Minga_-_01_-_Demonstration.mp3
     % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
   100 10550    0 10550    0     0  30403      0 --:--:-- --:--:-- --:--:-- 30403
   100 1025k  100 1025k    0     0  1847k      0 --:--:-- --:--:-- --:--:-- 1847k
   ```

1. Open <http://localhost:9000/mp3> in a browser or use `curl`:

   ```sh
   $ curl http://localhost:9000/mp3 --output Marsel_Minga_-_01_-_Demonstration.mp3
     % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
   100 1025k  100 1025k    0     0  71.5M      0 --:--:-- --:--:-- --:--:-- 71.5M
   ```
