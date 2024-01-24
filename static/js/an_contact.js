document.body.addEventListener('htmx:beforeRequest', function (event) {
    if (event.detail.elt === document.getElementById('contact-form')) {
        var sendingDiv = document.getElementById('sending-notification');
        var responseDiv = document.getElementById('contact-response');
        // Hide response div if it's visible
        if (responseDiv) {
            gsap.to(responseDiv, {
                opacity: 0,
                x: 20,
                duration: 0.5,
                onComplete: function () {
                    responseDiv.style.display = 'none';
                }
            });
        }
        document.getElementById('grid-first-name').value = '';
        document.getElementById('grid-last-name').value = '';
        document.getElementById('grid-email').value = '';
        document.getElementById('grid-message').value = '';

        // Show sending notification
        if (sendingDiv) {
            gsap.set(sendingDiv, {display: 'block', opacity: 0, x: 20});
            gsap.to(sendingDiv, {opacity: 1, x: 0, duration: 0.5});
        }
    }
});

document.body.addEventListener('htmx:afterOnLoad', function (event) {
    if (event.detail.elt === document.getElementById('contact-form')) {
        var responseDiv = document.getElementById('contact-response');
        var sendingDiv = document.getElementById('sending-notification');
        // Hide sending notification
        if (sendingDiv) {
            gsap.to(sendingDiv, {
                opacity: 0,
                x: 20,
                duration: 0.25,
                onComplete: function () {
                    sendingDiv.style.display = 'none';
                }
            });
        }

        // Show and animate the response div
        if (responseDiv) {
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