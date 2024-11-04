package abe

import {
	"rocprf"
}

// ABE.Setup(1λ) : It computes ppR← CPRF.Setup(1λ) and msk ← CPRF.KeyGen(pp), and outputs a public parameter pp and a master secret key msk.
func ABESetup() {

	msk, _ := KeyGen(modulus, length)
}

// ABE.KeyGen(msk, f) : It computes skf ← CPRF.Constrain(msk, f), and outputs skf .
func ABEKeyGen(msk *MasterKey) (skf *ConstrainedKey) {
	z := msk.z0
	csk, _ := msk.Constrain(z)
	skf := csk // not sure how this works, CEval has (csk *ConstrainedKey)
	return
}

// ABE.Enc(msk, x, msg) : It computes k ← CPRF.Eval(msk, x) and ct ← SKE.Enc(k, msg), and outputs ct
func ABEEnc(msk *MasterKey, x []*big.Int, m string) (ct string){ // is it string?
	k := msk.Eval(x)
	// ct := SKEEnc(k, m)
	// return
}

// ABE.Dec(skf , ct) : It computes k ← CPRF.CEval(skf , x), and outputs SKE.Dec(k, ct).
func ABEDec(csk *ConstrainedKey, x []*big.Int, ct string) (fm string){
	k = csk.CEval(x)
	// fm := SKEDec(k, ct)
	// return
}

