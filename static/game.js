var start = new Date().getTime();
var highScore = Number.MAX_SAFE_INTEGER;
var size = 500;

function makeShapeAppear() {
  var top = Math.random() * size;
  var left = Math.random() * size;
  var width = (Math.random() * size) / 2 + 100;
  if (Math.random() > 0.5) {
    document.getElementById("shape").style.borderRadius = "50%"; // Circle
  } else {
    document.getElementById("shape").style.borderRadius = "0"; // Square
  }
  document.getElementById("shape").style.width = width + "px";
  document.getElementById("shape").style.height = width + "px";

  document.getElementById("shape").style.top = top + "px";
  document.getElementById("shape").style.left = left + "px";

  document.getElementById("shape").style.display = "block";
  document.getElementById("shape").style.backgroundColor = getRandomColor();
  start = new Date().getTime();
}

function appearAfterDelay() {
  setTimeout(makeShapeAppear, Math.random() * 2000);
}

appearAfterDelay();
document.getElementById("shape").onclick = function () {
  document.getElementById("shape").style.display = "none";
  var end = new Date().getTime();
  var timeTaken = (end - start) / 1000;
  document.getElementById("timeTaken").innerHTML = timeTaken + "s";
  if (timeTaken < highScore) {
    highScore = timeTaken;
    document.getElementById("bestScore").innerHTML = highScore + "s";
  }
  appearAfterDelay();
};
function getRandomColor() {
  var letters = "0123456789ABCDEF";
  var color = "#";
  for (var i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
}
