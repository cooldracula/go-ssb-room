// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

let hasFocus = true;
window.addEventListener('blur', () => {
  hasFocus = false;
});
window.addEventListener('focus', () => {
  hasFocus = true;
});

const waitingElem = document.getElementById('waiting');
const failureElem = document.getElementById('failure');
const anchorElem = document.getElementById('alias-uri');

// Autoredirect to the ssb uri ASAP
setTimeout(() => {
  const ssbUri = anchorElem.href;
  window.location.replace(ssbUri);
  waitingElem.classList.remove('hidden');
  setTimeout(function () {
    if (hasFocus) {
      waitingElem.classList.add('hidden');
      failureElem.classList.remove('hidden');
    }
  }, 5000);
  window.location.replace(ssbUri);
}, 100);
