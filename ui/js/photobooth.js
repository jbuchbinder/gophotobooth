// photobooth.js - @jbuchbinder

var busy = false;
var debugEnabled = true;
var batch = '';
var delay = 2 * 1000; // 2 second delay

$(document).ready(function() {
	initDisplay();
	$(document).keypress(function(e) {
		if (e.which == 32) { // space
			startTakePhoto();
		}
	});
});

function initDisplay() {
	displayBig("PRESS BUTTON TO TAKE PICTURES");
}

function startTakePhoto() {
	if (busy) {
		console.log("startTakePhoto(): Already processing.");
		return;
	}
	debug("startTakePhoto()");
	busy = true;

	// Set batch name as date
	batch = new Date().toISOString();

	displayCountDown(function() {
		takePhoto(1, function() {
			displayCountDown(function() {
				takePhoto(2, function() {
					displayCountDown(function() {
						takePhoto(3, function() {
							displayCountDown(function() {
								takePhoto(4, function() {
									showThankYou();
									setTimeout(function() {
										initDisplay();
										busy = false;
									}, 3000);
								});
							});
						});
					});
				});
			});
		});
	});
}

function playAudio(sound) {
	$('#audio-'+sound)[0].play();
}

function displayBig(t) {
	$('#content').html(t);
	$('#text').bigtext();
	//$('body').append(t);
}

function displayCountDown(cb) {
	debug("displayCountDown()");
	playAudio("click");
	displayBig("- 3 -");
	setTimeout(function() {
		playAudio("click");
		displayBig("- 2 -");
		setTimeout(function() {
			playAudio("click");
			displayBig("- 1 -");
			setTimeout(function() {
				displayBig("CHEESE!");
				cb();
			}, delay);
		}, delay);
	}, delay);
}

function takePhoto(id, cb) {
	debug("takePhoto()");
	$.get("/api/photo/" + batch + "/" + id, function(data) {
		console.log("Returned : " + data);
	});
	playAudio("shutter");
	setTimeout(cb, delay * 2);
}

function showThankYou() {
	debug("showThankYou()");
	displayBig("THANK YOU");
}

function debug(s) {
	if (debugEnabled) {
		console.log(s);
	}
}

