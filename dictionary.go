package jiebago

import (
	"math"
	"sync"

	"github.com/rosbit/jiebago/dictionary"
)

// A Dictionary represents a thread-safe dictionary used for word segmentation.
type Dictionary struct {
	total, logTotal float64
	//freqMap         map[string]float64
	tokenAttrMap    map[string]dictionary.TokenAttr
	sync.RWMutex
}

// Load loads all tokens from given channel
func (d *Dictionary) Load(ch <-chan dictionary.Token) {
	d.Lock()
	for token := range ch {
		d.addToken(token)
	}
	d.Unlock()
	d.updateLogTotal()
}

// AddToken adds one token
func (d *Dictionary) AddToken(token dictionary.Token) {
	d.Lock()
	d.addToken(token)
	d.Unlock()
	d.updateLogTotal()
}

func (d *Dictionary) addToken(token dictionary.Token) {
	// d.freqMap[token.Text()] = token.Frequency()
	d.tokenAttrMap[token.Text()] = token.Attr()
	d.total += token.Frequency()
	runes := []rune(token.Text())
	n := len(runes)
	for i := 0; i < n; i++ { //TODO: n-1?
		frag := string(runes[:i+1])
		// if _, ok := d.freqMap[frag]; !ok {
			// d.freqMap[frag] = 0.0
		if _, ok := d.tokenAttrMap[frag]; !ok {
			d.tokenAttrMap[frag] = dictionary.TokenAttr{}
		}
	}
}

func (d *Dictionary) updateLogTotal() {
	d.logTotal = math.Log(d.total)
}

// Frequency returns the frequency and existence of give word
func (d *Dictionary) Frequency(key string) (float64, bool) {
	d.RLock()
	defer d.RUnlock()
	// freq, ok := d.freqMap[key]
	// return freq, ok
	if attr, ok := d.tokenAttrMap[key]; ok {
		return attr.Frequency(), ok
	}
	return 0.0, false
}

func (d *Dictionary) Attr(key string) (float64, string, bool) {
	d.RLock()
	defer d.RUnlock()
	if attr, ok := d.tokenAttrMap[key]; ok {
		return attr.Frequency(), attr.Pos(), true
	}
	return 0.0, "", false
}

func (d *Dictionary) loadDictionary(fileName string) error {
	return dictionary.LoadDictionary(d, fileName)
}
