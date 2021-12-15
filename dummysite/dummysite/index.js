const scrape = require("website-scraper")
const express = require("express")
const PORT = 3000
const websiteUrl = process.env.WEBSITE_URL || "https://example.com"

const options = {
    urls: [websiteUrl],
    directory: "dummysite/"
}

const app = express()

scrape(options).then(() => {
    app.use(express.static("dummysite"))
    console.log(`Fetched website ${websiteUrl}`)
    app.listen(PORT, console.log(`Server listening on port: ${PORT}`))
})