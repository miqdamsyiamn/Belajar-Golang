// Create binary animation background
document.addEventListener('DOMContentLoaded', function() {
	const binaryBg = document.getElementById('binary-bg');
	const width = window.innerWidth;
	const height = window.innerHeight;
	const columns = Math.floor(width / 20);
	
	// Create binary streams
	for (let i = 0; i < columns; i++) {
		const stream = document.createElement('div');
		stream.className = 'binary-stream';
		stream.style.left = (i * 20) + 'px';
		stream.style.animationDelay = (Math.random() * 2) + 's';
		
		// Generate random binary sequence
		let binaryText = '';
		for (let j = 0; j < 50; j++) {
			binaryText += Math.round(Math.random());
		}
		stream.textContent = binaryText;
		binaryBg.appendChild(stream);
	}
});