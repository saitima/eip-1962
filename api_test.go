package eip

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestG1MulPoint(t *testing.T) {
	file := "test_vectors/custom/256.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeG1MulBinary()
	if err != nil {
		t.Fatal(err)
	}

	api := new(g1Api)
	actual, err := api.mulPoint(in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
	}
}
func TestG22MulPoint(t *testing.T) {
	file := "test_vectors/custom/256.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeG22MulBinary()
	if err != nil {
		t.Fatal(err)
	}

	api := new(API)
	actual, err := api.Run(OPERATION_G2_MUL, in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}
func TestG23MulPoint(t *testing.T) {
	file := "test_vectors/custom/320_cubic.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeG23MulBinary()
	if err != nil {
		t.Fatal(err)
	}

	api := new(API)
	actual, err := api.Run(OPERATION_G2_MUL, in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}

func TestBNPairingApi(t *testing.T) {
	file := "test_vectors/custom/256.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeBNPairingBinary()
	if err != nil {
		t.Fatal(err)
	}
	actual, err := new(API).Run(OPERATION_BNPAIR, in)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}
func TestBLSPairingApi(t *testing.T) {
	file := "test_vectors/custom/384.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeBLSPairingBinary()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := new(API).Run(OPERATION_BLS12PAIR, in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}

func TestMNT4Pairing(t *testing.T) {
	file := "test_vectors/custom/320.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeMNT4PairingBinary()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := new(API).Run(OPERATION_MNT4PAIR, in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}

func TestMNT4753PairingApi(t *testing.T) {
	in, err := hex.DecodeString("5f01c4c62d92c41110229022eee2cdadb7f997505b8fafed5eb7e8f96c97d87307fdb925e8a0ed8d99d124d9a15af79db117e776f218059db80f0da5cb537e38685acce9767254a4638810719ac425f0e39d54522cdd119f5e9063de245e8001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000201373684a8c9dcae7a016ac5d7748d3313cd8e39051c596560835df0c9e50a5b59b882a92c78dc537e51a16703ec9855c77fc3d8bb21c8d68bb8cfb9db4b8c8fba773111c36c8b1b4e8f1ece940ef9eaad265458e06372009c9a0491678ef4600001c4c62d92c41110229022eee2cdadb7f997505b8fafed5eb7e8f96c97d87307fdb925e8a0ed8d99d124d9a15af79db26c5c28c859a99b3eebca9429212636b9dff97634993aa4d6c381bc3f0057974ea099170fa13a4fd90776e240000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d30015474b1d641a3fd86dcbcee5dcda7fe51852c8cbe26e600733b714aa43c31a66b0344c4e2c428b07a7713041ba180000130015474b1d641a3fd86dcbcee5dcda7fe51852c8cbe26e600733b714aa43c31a66b0344c4e2c428b07a7713041ba17fff0101010401013b42397c8b004d06f0e98fbc12e8ee65adefcdba683c5630e6b58fb69610b02eab1d43484ddfab28213098b562d799243fb14330903aa64878cfeb34a45d1285da665f5c3f37eb76b86209dcd081ccaef03e65f33d490de480bfee06db00e3eb479d308664381e7942d6c522c0833f674296169420f1dd90680d0ba6686fc27549d52e4292ea5d611cb6b0df32545b07f281032d0a71f8d485e6907766462e17e8dd55a875bd36fe4cd42cac31c0629fb26c333fe091211d0561d10e00f1b7155ed4e903332835a5de0f327aa11b2d74eb8627e3a7b833be42c11d044b5cf0ae49850eeb07d90c77c67256474b2febf924aca0bfa2e4dacb821c91a04fd0165ac8debb2fc1e763a5c32c2c9f572caa85a91c5243ec4b2981af890400d49c264ec663e731713182a88907b8e979ced82ca592777ad052ec5f4b95dc78dc2010d74f82b9e6d066813ed67f3af1de0d5d425da7a19916cf103f102adf5f95b6b62c24c7d186d60b4a103e157e5667038bb2e828a3374d6439526272004b0e2fef08096ebbaddd2d7f288c4acf17b2267e21dc5ce0f925cd5d02209e34d8b69cc94aef5d90af34d3cd98287ace8f1162079cd2d3d7e6c6c2c073c24a359437e75638a1458f4b2face11f8d2a5200b14d6f9dd0fdd407f04be620ee00bc1925e7fcb64f6f8697cd5e45fae22f5688e51b30bd984c0acdc67d2962520e80d31966e3ec477909ecca358be2eee53c75f55a6f7d9660dd6f3d4336ad50e8bfa5375791d73b863d59c422c3ea006b013e7afb186f2eaa9df68f4d609801013b42397c8b004d06f0e98fbc12e8ee65adefcdba683c5630e6b58fb69610b02eab1d43484ddfab28213098b562d799243fb14330903aa64878cfeb34a45d1285da665f5c3f37eb76b86209dcd081ccaef03e65f33d490de480bfee06db00e0dae5f5938aabea71a9ac0c088af77657e918f999593dc60b69048acccc9f8df6b09ecbbf4b06e6c77884a446be7ec38c6eff970270ad9d14d1456cedc102149ed18d94fefbedcad9734deff944b1dcf1b27a70de5f7dff42c11efcaef300f1b7155ed4e903332835a5de0f327aa11b2d74eb8627e3a7b833be42c11d044b5cf0ae49850eeb07d90c77c67256474b2febf924aca0bfa2e4dacb821c91a04fd0165ac8debb2fc1e763a5c32c2c9f572caa85a91c5243ec4b2981af890400d49c264ec663e731713182a88907b8e979ced82ca592777ad052ec5f4b95dc78dc2010d74f82b9e6d066813ed67f3af1de0d5d425da7a19916cf103f102adf5f95b6b62c24c7d186d60b4a103e157e5667038bb2e828a3374d6439526272004b0e2fef08096ebbaddd2d7f288c4acf17b2267e21dc5ce0f925cd5d02209e34d8b69cc94aef5d90af34d3cd98287ace8f1162079cd2d3d7e6c6c2c073c24a359437e75638a1458f4b2face11f8d2a5200b14d6f9dd0fdd407f04be620ee00bc1925e7fcb64f6f8697cd5e45fae22f5688e51b30bd984c0acdc67d2962520e80d31966e3ec477909ecca358be2eee53c75f55a6f7d9660dd6f3d4336ad50e8bfa5375791d73b863d59c422c3ea006b013e7afb186f2eaa9df68f4d609801013b42397c8b004d06f0e98fbc12e8ee65adefcdba683c5630e6b58fb69610b02eab1d43484ddfab28213098b562d799243fb14330903aa64878cfeb34a45d1285da665f5c3f37eb76b86209dcd081ccaef03e65f33d490de480bfee06db00e3eb479d308664381e7942d6c522c0833f674296169420f1dd90680d0ba6686fc27549d52e4292ea5d611cb6b0df32545b07f281032d0a71f8d485e6907766462e17e8dd55a875bd36fe4cd42cac31c0629fb26c333fe091211d0561d10e00f1b7155ed4e903332835a5de0f327aa11b2d74eb8627e3a7b833be42c11d044b5cf0ae49850eeb07d90c77c67256474b2febf924aca0bfa2e4dacb821c91a04fd0165ac8debb2fc1e763a5c32c2c9f572caa85a91c5243ec4b2981af890400d49c264ec663e731713182a88907b8e979ced82ca592777ad052ec5f4b95dc78dc2010d74f82b9e6d066813ed67f3af1de0d5d425da7a19916cf103f102adf5f95b6b62c24c7d186d60b4a103e157e5667038bb2e828a3374d6439526272004b0e2fef08096ebbaddd2d7f288c4acf17b2267e21dc5ce0f925cd5d02209e34d8b69cc94aef5d90af34d3cd98287ace8f1162079cd2d3d7e6c6c2c073c24a359437e75638a1458f4b2face11f8d2a5200b14d6f9dd0fdd407f04be620ee00bc1925e7fcb64f6f8697cd5e45fae22f5688e51b30bd984c0acdc67d2962520e80d31966e3ec477909ecca358be2eee53c75f55a6f7d9660dd6f3d4336ad50e8bfa5375791d73b863d59c422c3ea006b013e7afb186f2eaa9df68f4d609801013b42397c8b004d06f0e98fbc12e8ee65adefcdba683c5630e6b58fb69610b02eab1d43484ddfab28213098b562d799243fb14330903aa64878cfeb34a45d1285da665f5c3f37eb76b86209dcd081ccaef03e65f33d490de480bfee06db00e0dae5f5938aabea71a9ac0c088af77657e918f999593dc60b69048acccc9f8df6b09ecbbf4b06e6c77884a446be7ec38c6eff970270ad9d14d1456cedc102149ed18d94fefbedcad9734deff944b1dcf1b27a70de5f7dff42c11efcaef300f1b7155ed4e903332835a5de0f327aa11b2d74eb8627e3a7b833be42c11d044b5cf0ae49850eeb07d90c77c67256474b2febf924aca0bfa2e4dacb821c91a04fd0165ac8debb2fc1e763a5c32c2c9f572caa85a91c5243ec4b2981af890400d49c264ec663e731713182a88907b8e979ced82ca592777ad052ec5f4b95dc78dc2010d74f82b9e6d066813ed67f3af1de0d5d425da7a19916cf103f102adf5f95b6b62c24c7d186d60b4a103e157e5667038bb2e828a3374d6439526272004b0e2fef08096ebbaddd2d7f288c4acf17b2267e21dc5ce0f925cd5d02209e34d8b69cc94aef5d90af34d3cd98287ace8f1162079cd2d3d7e6c6c2c073c24a359437e75638a1458f4b2face11f8d2a5200b14d6f9dd0fdd407f04be620ee00bc1925e7fcb64f6f8697cd5e45fae22f5688e51b30bd984c0acdc67d2962520e80d31966e3ec477909ecca358be2eee53c75f55a6f7d9660dd6f3d4336ad50e8bfa5375791d73b863d59c422c3ea006b013e7afb186f2eaa9df68f4d6098")
	if err != nil {
		t.Fatal(err)
	}
	_, err = new(API).Run(OPERATION_MNT4PAIR, in)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMNT6Pairing(t *testing.T) {
	file := "test_vectors/custom/320_mnt6.json"
	v, err := newTestVectorJSONFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	in, expected, err := v.makeMNT6PairingBinary()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := new(API).Run(OPERATION_MNT6PAIR, in)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Logf("actual %x\n", actual)
		t.Logf("expected %x\n", expected)
		t.Fatalf("not equal")
	}
}
