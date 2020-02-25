function bytesToHexString(byteArray) {
    return Array.from(byteArray, function(byte) {
      return ('0' + (byte & 0xFF).toString(16)).slice(-2);
    }).join('')
}

function hexStringToBytes(hex) {
  var result = [];
  while (hex.length >= 8) { 
      result.push(parseInt(hex.substring(0, 8), 16));
  
      hex = hex.substring(8, hex.length);
  }

  return result;
}

function intToBytes(x) {
  let y= Math.floor(x/2**32);
  return [y,(y<<8),(y<<16),(y<<24), x,(x<<8),(x<<16),(x<<24)].map(z=> z>>>24)
}

function utf8ToBytes(str) {
  var utf8 = [];
  for (var i=0; i < str.length; i++) {
      var charcode = str.charCodeAt(i);
      if (charcode < 0x80) utf8.push(charcode);
      else if (charcode < 0x800) {
          utf8.push(0xc0 | (charcode >> 6), 
                    0x80 | (charcode & 0x3f));
      }
      else if (charcode < 0xd800 || charcode >= 0xe000) {
          utf8.push(0xe0 | (charcode >> 12), 
                    0x80 | ((charcode>>6) & 0x3f), 
                    0x80 | (charcode & 0x3f));
      }
      // surrogate pair
      else {
          i++;
          // UTF-16 encodes 0x10000-0x10FFFF by
          // subtracting 0x10000 and splitting the
          // 20 bits of 0x0-0xFFFFF into two halves
          charcode = 0x10000 + (((charcode & 0x3ff)<<10)
                    | (str.charCodeAt(i) & 0x3ff));
          utf8.push(0xf0 | (charcode >>18), 
                    0x80 | ((charcode>>12) & 0x3f), 
                    0x80 | ((charcode>>6) & 0x3f), 
                    0x80 | (charcode & 0x3f));
      }
  }
  return utf8;
}


function splitSign(sig) {
  let r = sig.slice(0, 66)
  let s = "0x"+sig.slice(66, 130)
  let v = parseInt("0x"+sig.slice(130, 132)) + 27
  return {
    r: r,
    s: s,
    v: v
  };
}

module.exports.bytesToHexString = bytesToHexString;
module.exports.hexStringToBytes = hexStringToBytes;
module.exports.intToBytes = intToBytes;
module.exports.utf8ToBytes = utf8ToBytes;
module.exports.splitSign = splitSign;