var urls = [];
var divs = document.querySelectorAll("div._mck9w._gvoze._tn0ps");
for (var i = 0; i < divs.length; i++) {
  var div = divs[i];
  urls.push(div.firstChild["href"]);
}
console.log(urls);
