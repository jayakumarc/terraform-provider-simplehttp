
data "simplehttp" "google_page" {
    url = "https://www.google.com/"
}

output "content" {
  value = "${data.simplehttp.google_page.body}"
}