window.addEventListener('load', function () {
    gsap.registerPlugin(ScrollTrigger);

    gsap.from('.project-title', {
        duration: 1,
        y: -50,
        opacity: 0,
        ease: 'power1.out',
        scrollTrigger: {
            trigger: '.project-title',
            start: 'top 90%',
            end: 'bottom 60%',
            toggleActions: 'play none none none'
        }
    });

    gsap.from('.project-card', {
        duration: 1,
        x: -200,
        opacity: 0,
        ease: 'power1.out',
        scrollTrigger: {
            trigger: '.project-card',
            start: 'top 90%',
            end: 'bottom 60%',
            toggleActions: 'play none none none'
        }
    });

    gsap.from('.project-image', {
        duration: 1.5,
        scale: 0.5,
        opacity: 0,
        ease: 'power1.out',
        scrollTrigger: {
            trigger: '.project-image',
            start: 'top 80%',
            end: 'bottom 60%',
            toggleActions: 'play none none none'
        }
    });

    gsap.from('.project-description', {
        duration: 1,
        x: 200,
        opacity: 0,
        ease: 'power1.out',
        scrollTrigger: {
            trigger: '.project-description',
            start: 'top 90%',
            end: 'bottom 60%',
            toggleActions: 'play none none none'
        }
    });

});