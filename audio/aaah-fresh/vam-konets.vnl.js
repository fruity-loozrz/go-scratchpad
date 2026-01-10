sample('./audio/aaah-fresh/vam-konets.wav').offset(0.1);

bpm(100);
rpm(33);

function pauseInBeats(beats) {
    $()
        .platterEnvelopInBeats(from(0, 0).to('Linear', beats, 0))
        .faderEnvelope(micro('_', 'OutExpo'));
}

function wreke() {
    $()
        .platterEnvelopInBeats(from(0, 0).to('InOutSine', 0.5, 0.3))
        .faderEnvelope(micro('---', 'InOutExpo'));
}

function woh() {
    $()
        .platterEnvelopInBeats(from(0, 0).to('OutCubic', 1, 0.1))
        .faderEnvelope(micro('--___', 'InSine'));
}

function wuee() {
    $()
        .platterEnvelopInBeats(from(0, 0).to('InOutBack', 0.5, 0.09))
        .faderEnvelope(micro('-----', 'InSine'));
}

wreke()
woh()
wreke()
wreke()
wuee()
woh()

pauseInBeats(0.125)

wreke()
woh()
wreke()
wreke()
wuee()
woh()
