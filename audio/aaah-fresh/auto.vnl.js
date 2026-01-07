sample('./audio/aaah-fresh/aaah-fresh.wav');
bpm(100);
rpm(33);
seed(1);

$().Dur(0.25).Easing('heavy').Platter(at(0), at(0.1)).FaderPattern('open');
$().Dur(0.25).Easing('heavy').Platter(at(0.1), at(0)).FaderPattern('open');
$().Dur(0.25).Easing('heavy').Platter(at(0), at(0.1)).FaderPattern('open');
$().Dur(0.25).Easing('heavy').Platter(at(0.1), at(0)).FaderPattern('open');
$().Dur(0.5).Easing('smooth').Platter(at(0), at(0.4)).FaderPattern('open');
$().Dur(0.25).Easing('heavy').Platter(at(0), at(0.1)).FaderPattern('open');
$().Dur(0.25).Easing('heavy').Platter(at(0.1), at(0)).FaderPattern('open');
$().Dur(0.5).Easing('bounce').Platter(at(0), at(0.4)).FaderPattern('open');
$().Dur(0.5).Easing('bounce').Platter(at(0), at(0.4)).FaderPattern('open');
$().Dur(0.5).Easing('smooth').Platter(at(0.3), at(0.2)).FaderPattern('flare2');

function at(t) {
    const start = 0.0 + rand() * 0.03;
    const rotationSpeed = 0.25 + rand() * 0.01;
    return start + t * rotationSpeed;
}
