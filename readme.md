# App Store Reviewer
There are 2 projects in this repo, a Go REST Api in `/api` and a React app in `/app`

The `/api` directory contains the Go api, which runs in port 8080. In order to run, just `cd` into the directory and run `go run main.go`. 

The `/app` directory contains the React app, which runs in port 3000. On first download, follow these steps:
1. `npx install`
2. `npm run build`
3. `npm run start`

## Adding more apps

The file `./api/application.json` contains an JSON array with application ids. In order to add or change apps to the monitor, just add the app's ID to the JSON array, and save the file. On the next poll, the corresponding file will be created, and the app will be available to query.

## External libraries

The API uses Gin framework for REST API. To schedule the cron job that polls the appstore api, the cron v3 library was used. From the Gin framework, the CORS library was also imported. 
The react app url as well as the frecuency of the cron is set directly in the code, it should be set in environment variables or similar.

For the react app, Next.js and tailwind.css were used as I'm more familiar with them to work with React.
