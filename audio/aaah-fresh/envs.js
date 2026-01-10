sample('./audio/aaah-fresh/aaah-fresh.wav');
bpm(100);
rpm(33);

$()
    .platterEnvelopInBeats(
        from(0, 0)
            .to('InSine', 0.25, 0.2)
            .to('InOutSine', 0.5, 0.1)
        )
    .faderEnvelope(micro('__-_-_-__', 'OutExpo'));

$()
    .platterEnvelopInBeats(
        from(0, 0)
            .to('InSine', 0.25, 0.2)
            .to('InOutSine', 0.5, 0.1)
        )
    .faderEnvelope(micro('__-___--', 'OutExpo'));


$()
    .platterEnvelopInBeats(
        from(0, 0)
            .to('InOutSine', 1, 0.2)
        )
    .faderEnvelope(micro('_-_-_-_-_-_-_', 'OutExpo'));
