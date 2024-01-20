document.getElementById('rateForm').addEventListener
('submit', function (event) {
    event.preventDefault();

    const cloudHours = parseFloat(document.getElementById('cloudServiceHours').value) || 0;
    const ciCdHours = parseFloat(document.getElementById('ciCdHours').value) || 0;
    const monitoringHours = parseFloat(document.getElementById('monitoringHours').value) || 0;
    const automationHours = parseFloat(document.getElementById('automationHours').value) || 0;

    if (cloudHours <= 0 && ciCdHours <= 0 && monitoringHours <= 0 && automationHours <= 0) {
        htmx.ajax('GET', '/calculator-popup', '#popupContainer');
        return;
    }

    const cloudRate = 50;
    const ciCdRate = 40;
    const monitoringRate = 45;
    const automationRate = 55;


    const totalCost = (cloudHours * cloudRate) + (ciCdHours * ciCdRate) + (monitoringHours * monitoringRate) + (automationHours * automationRate);
    document.getElementById('totalCost').textContent = `Total Cost: $${totalCost.toFixed(2)}`;
});

document.body.addEventListener('htmx:afterOnLoad', function () {
    var closeModalButton = document.getElementById('closeModal');
    if (closeModalButton) {
        closeModalButton.addEventListener('click', function () {
            document.getElementById('popupModal').classList.add('hidden');
        });
    }
});

window.addEventListener('load', function () {
    gsap.registerPlugin(ScrollTrigger);

    gsap.from('.calculator-title', {
        opacity: 0,
        x: -50,
        duration: 2,
        ease: 'power2.out',
        scrollTrigger: {
            trigger: '.calculator-title',
            start: 'top bottom',
            toggleActions: 'play none none none'
        }
    });

    gsap.from('.calculator-container', {
        opacity: 0,
        x: 50,
        duration: 2,
        ease: 'power2.out',
        scrollTrigger: {
            trigger: '.calculator-container',
            start: 'top bottom',
            toggleActions: 'play none none none'
        }
    });

    ScrollTrigger.refresh();
});