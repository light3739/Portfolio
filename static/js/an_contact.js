// Message
document.body.addEventListener('htmx:afterOnLoad', function (event) {
    if (event.detail.elt === document.getElementById('contact-form')) {
        // Your code for handling the event triggered by the contact form
        var responseDiv = document.getElementById('contact-response');
        if (responseDiv) {
            // Show and animate the response div
            gsap.set(responseDiv, {display: 'block', opacity: 0, x: 20});
            gsap.to(responseDiv, {opacity: 1, x: 0, duration: 0.5});

            // Hide the message after 5 seconds with animation
            setTimeout(function () {
                gsap.to(responseDiv, {
                    opacity: 0,
                    x: 20,
                    duration: 0.5,
                    onComplete: function () {
                        responseDiv.style.display = 'none';
                    }
                });
            }, 3000);
        }
    }
});



gsap.utils.toArray('input, textarea').forEach(field => {
    field.addEventListener('focus', () => {
        gsap.to(field, {borderColor: '#4A90E2', borderWidth: '2px', duration: 0.3});
    });
    field.addEventListener('blur', () => {
        gsap.to(field, {borderColor: '#d1d5db', borderWidth: '1px', duration: 0.3});
    });
});

// Button Hover Effect
const sendButton = document.querySelector('#contactSendButton');
if (sendButton) {
    sendButton.addEventListener('mouseenter', () => {
        gsap.to(sendButton, {scale: 1.10, duration: 0.5, repeat: -1, yoyo: true});
    });
    sendButton.addEventListener('mouseleave', () => {
        gsap.killTweensOf(sendButton);
        gsap.to(sendButton, {scale: 1, duration: 0.5});
    });
}

// Background Animation
gsap.to('body', {backgroundPosition: '100% 100%', ease: 'none', duration: 200, repeat: -1, yoyo: true});