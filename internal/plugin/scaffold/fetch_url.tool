tool "fetch_url" {
  description = "Fetch a URL and return the page contents as markdown"

  param "url" {
    type        = string
    description = "The absolute URL to fetch"
  }

  returns {
    type = string
  }

  source {
    kind = "mcp"
    uri  = "mcp://fetch/fetch"
  }
}
