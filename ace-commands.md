# Ace Editor Commands

session = editor.getSession();
doc = session.getDocument();
lines = doc.$lines;

mode = session.getMode();

keyboardHandler = editor.getKeyboardHandler();
editor.setKeyboardHandler('ace/keyboard/vim');

editor.on('change', (delta) => {
    sendUpdate(delta);
});

doc.applyDeta(delta);

