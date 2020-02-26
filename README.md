# OriginStamp Proofs

OriginStamp returns proofs

The PDF variant includes an explanation of what's need to verify the proof:

1. Determine the SHA-256 of your data
1. Validate your proof
1. Determine the private key
1. Determine the bitcoin address
1. Check the transactions

This sample implements the proof in Golang.

It's best experienced through the tests but there's a `main.go` for giggles.

Assuming we start with such a proof:

```XML
<node value="69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950" type="key">
	<left value="e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e" type="mesh">
		<left value="ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388" type="mesh">
			<left value="9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808" type="mesh">
				<left value="b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f" type="mesh">
					<left value="4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e" type="mesh">
						<left value="10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5" type="mesh">
							<left value="c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d" type="mesh">
								<left value="c7ebc60a38854a4671ea786f14ceff0e7e183f75d6fbda3557292bae3b13f23d" type="mesh"/>
								<right value="a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4" type="mesh">
									<left value="d1776b7c1b51ad1441bf765e194377ac9060f2402d6e91cee5502d8ca3a99546" type="mesh"/>
									<right value="25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1" type="mesh">
										<left value="c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae" type="mesh">
											<left value="097c397d862795e96abb576eb6bc76365f0235cfa1c938352045d1e62b347b1d" type="mesh"/>
											<right value="9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb" type="mesh">
												<left value="369cb74a155479eecf22f65337e03a4ad5b935772e2d3348854f3d9d71923960" type="mesh"/>
												<right value="e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b" type="mesh">
													<left value="0d899a0e8f72691e3859377c4780a5b54c40e743d1486aeb1e76c038dab9e8ce" type="mesh"/>
													<right value="e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7" type="mesh">
														<left value="037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b" type="hash"/>
														<right value="038df16cfbbeed2f029f00587bfcaa9920d4f7bcaeee60ee788d4a66179709c3" type="mesh"/>
													</right>
												</right>
											</right>
										</left>
										<right value="c3d44098be84f892ee1fe1c2c6caa52257eb96caaaa3f313edbafdc4a3ba6eae" type="mesh"/>
									</right>
								</right>
							</left>
							<right value="5b2e60ccb6907fae79d249bf7943b050c28b5fef84b9b44399f58461dd120c4a" type="mesh"/>
						</left>
						<right value="c47e581dd90c10aad2cb4121c1a94a3199f9788f4f2e5d0f200f8f167ae2517c" type="mesh"/>
					</left>
					<right value="dbdfc7af18a876d64d6fa2f0110e5dc6ead648f3e33f9e0163ceb267d47cd90d" type="mesh"/>
				</left>
				<right value="98f7dafd1f8c6ec84a555819d9cd05baed657e3973f86f3fb178dbd13c676bba" type="mesh"/>
			</left>
			<right value="36447318383768a406815764e128d404ca1482ec57cd995e12ac97962b33b73e" type="mesh"/>
		</left>
		<right value="6e7ca4d2ea8919396dfd66a7d57338cb0ce8bca76043b5a4745e463c2c9d9af8" type="mesh"/>
	</left>
	<right value="0f1301c53b6885f23037a81da41f450968a2b133f28b30d65d1fb431da74a9c1" type="mesh"/>
</node>
```

### Tree: Value()

We marshal this into a `Tree` and then calculate its `Value()`

This will error if the tree's `node` value does not correctly reflect the aggregation of its children's hashes.


### Bitcoin: PublicKey(), Address()

We then determine the `PublicKey()` of this `node` value (i.e. the private key)

I needed help from some friends: https://stackoverflow.com/q/60385053/609290

I had 2 primary issues:

+ new(ing) an `ECDSA.PrivateKey` from a hex-encoded private key
+ Identifying that the common default curve ([P256](https://golang.org/pkg/crypto/elliptic/#P256) aka secp256r1) was incorrect and that I need to use secp256k1 (and find a Golang implementation)

I learned more about Elliptic Curves too which is always a good thing, including some nitty-gritty on [the difference between secp256r1 and secp256k1](https://www.johndcook.com/blog/2018/08/21/a-tale-of-two-elliptic-curves/)

> Ethereum has a robust Golang implementation and, interestingly (!?) Ethereum and Bitcoin both use secp256k1, so I used Ethereum's [secp256k1](https://godoc.org/github.com/ethereum/go-ethereum/crypto/secp256k1) module

Given a secpk256k1 hex-encoded private key , the following Golang function will return the hex-encoded public key:

```Golang
func Public(privateKey string) (publicKey string) {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return fmt.Sprintf("%x", elliptic.Marshal(secp256k1.S256(), e.X, e.Y))
}
```

Finally, we compute the Bitcoin `Address` following the ordained steps and implemented by [TP's Bitcoin Address Tests](https://gobittest.appspot.com/Address)

Lastly, we check a Bitcoin transaction tool for the address in the transaction. It should be the first.

## Golang

```Golang
go test ./...
ok  	github.com/DazWilkin/proof	0.017s
```

Or:

```Golang
go run ./...
```

yield:

```
go run ./...
2020/02/26 14:07:59  <node value="69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950" type="key">
    <left value="e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e" type="mesh">
       <left value="ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388" type="mesh">
          <left value="9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808" type="mesh">
             <left value="b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f" type="mesh">
                <left value="4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e" type="mesh">
                   <left value="10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5" type="mesh">
                      <left value="c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d" type="mesh">
                         <left value="c7ebc60a38854a4671ea786f14ceff0e7e183f75d6fbda3557292bae3b13f23d" type="mesh"></left>
                         <right value="a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4" type="mesh">
                            <left value="d1776b7c1b51ad1441bf765e194377ac9060f2402d6e91cee5502d8ca3a99546" type="mesh"></left>
                            <right value="25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1" type="mesh">
                               <left value="c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae" type="mesh">
                                  <left value="097c397d862795e96abb576eb6bc76365f0235cfa1c938352045d1e62b347b1d" type="mesh"></left>
                                  <right value="9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb" type="mesh">
                                     <left value="369cb74a155479eecf22f65337e03a4ad5b935772e2d3348854f3d9d71923960" type="mesh"></left>
                                     <right value="e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b" type="mesh">
                                        <left value="0d899a0e8f72691e3859377c4780a5b54c40e743d1486aeb1e76c038dab9e8ce" type="mesh"></left>
                                        <right value="e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7" type="mesh">
                                           <left value="037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b" type="hash"></left>
                                           <right value="038df16cfbbeed2f029f00587bfcaa9920d4f7bcaeee60ee788d4a66179709c3" type="mesh"></right>
                                        </right>
                                     </right>
                                  </right>
                               </left>
                               <right value="c3d44098be84f892ee1fe1c2c6caa52257eb96caaaa3f313edbafdc4a3ba6eae" type="mesh"></right>
                            </right>
                         </right>
                      </left>
                      <right value="5b2e60ccb6907fae79d249bf7943b050c28b5fef84b9b44399f58461dd120c4a" type="mesh"></right>
                   </left>
                   <right value="c47e581dd90c10aad2cb4121c1a94a3199f9788f4f2e5d0f200f8f167ae2517c" type="mesh"></right>
                </left>
                <right value="dbdfc7af18a876d64d6fa2f0110e5dc6ead648f3e33f9e0163ceb267d47cd90d" type="mesh"></right>
             </left>
             <right value="98f7dafd1f8c6ec84a555819d9cd05baed657e3973f86f3fb178dbd13c676bba" type="mesh"></right>
          </left>
          <right value="36447318383768a406815764e128d404ca1482ec57cd995e12ac97962b33b73e" type="mesh"></right>
       </left>
       <right value="6e7ca4d2ea8919396dfd66a7d57338cb0ce8bca76043b5a4745e463c2c9d9af8" type="mesh"></right>
    </left>
    <right value="0f1301c53b6885f23037a81da41f450968a2b133f28b30d65d1fb431da74a9c1" type="mesh"></right>
 </node>
2020/02/26 14:07:59 Value: [69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950]
2020/02/26 14:07:59 Value: [e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e]
2020/02/26 14:07:59 Value: [ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388]
2020/02/26 14:07:59 Value: [9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808]
2020/02/26 14:07:59 Value: [b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f]
2020/02/26 14:07:59 Value: [4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e]
2020/02/26 14:07:59 Value: [10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5]
2020/02/26 14:07:59 Value: [c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d]
2020/02/26 14:07:59 Value: [c7ebc60a38854a4671ea786f14ceff0e7e183f75d6fbda3557292bae3b13f23d]
2020/02/26 14:07:59 Value: [c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d]  L=c7ebc60a38854a4671ea786f14ceff0e7e183f75d6fbda3557292bae3b13f23d
2020/02/26 14:07:59 Value: [a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4]
2020/02/26 14:07:59 Value: [d1776b7c1b51ad1441bf765e194377ac9060f2402d6e91cee5502d8ca3a99546]
2020/02/26 14:07:59 Value: [a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4]  L=d1776b7c1b51ad1441bf765e194377ac9060f2402d6e91cee5502d8ca3a99546
2020/02/26 14:07:59 Value: [25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1]
2020/02/26 14:07:59 Value: [c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae]
2020/02/26 14:07:59 Value: [097c397d862795e96abb576eb6bc76365f0235cfa1c938352045d1e62b347b1d]
2020/02/26 14:07:59 Value: [c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae]  L=097c397d862795e96abb576eb6bc76365f0235cfa1c938352045d1e62b347b1d
2020/02/26 14:07:59 Value: [9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb]
2020/02/26 14:07:59 Value: [369cb74a155479eecf22f65337e03a4ad5b935772e2d3348854f3d9d71923960]
2020/02/26 14:07:59 Value: [9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb]  L=369cb74a155479eecf22f65337e03a4ad5b935772e2d3348854f3d9d71923960
2020/02/26 14:07:59 Value: [e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b]
2020/02/26 14:07:59 Value: [0d899a0e8f72691e3859377c4780a5b54c40e743d1486aeb1e76c038dab9e8ce]
2020/02/26 14:07:59 Value: [e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b]  L=0d899a0e8f72691e3859377c4780a5b54c40e743d1486aeb1e76c038dab9e8ce
2020/02/26 14:07:59 Value: [e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7]
2020/02/26 14:07:59 Value: [037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b]
2020/02/26 14:07:59 Value: [e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7]  L=037e945cf8da5945acbcf2390c71a497c6edefdc364ada1f33d76a2b5f8b472b
2020/02/26 14:07:59 Value: [038df16cfbbeed2f029f00587bfcaa9920d4f7bcaeee60ee788d4a66179709c3]
2020/02/26 14:07:59 Value: [e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7]  R=038df16cfbbeed2f029f00587bfcaa9920d4f7bcaeee60ee788d4a66179709c3
2020/02/26 14:07:59 Value: [e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7] LR=e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7
2020/02/26 14:07:59 Value: [e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b]  R=e40a25b8bd05b48675a02c8ad7e0c23b6d62385b0acaffbab63f95e39369ebc7
2020/02/26 14:07:59 Value: [e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b] LR=e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b
2020/02/26 14:07:59 Value: [9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb]  R=e220f15411b706ad10ff9990dc7f760ddfec2ebc61a11204fc5b36e140c17e0b
2020/02/26 14:07:59 Value: [9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb] LR=9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb
2020/02/26 14:07:59 Value: [c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae]  R=9407b62021fe6335430fc54090b90e6588ac38421309e279460335b5f94d67fb
2020/02/26 14:07:59 Value: [c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae] LR=c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae
2020/02/26 14:07:59 Value: [25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1]  L=c006e869d7ba65ac4d5105888fee6a6bf6a3b38f23c6300137f003a3988238ae
2020/02/26 14:07:59 Value: [c3d44098be84f892ee1fe1c2c6caa52257eb96caaaa3f313edbafdc4a3ba6eae]
2020/02/26 14:07:59 Value: [25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1]  R=c3d44098be84f892ee1fe1c2c6caa52257eb96caaaa3f313edbafdc4a3ba6eae
2020/02/26 14:07:59 Value: [25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1] LR=25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1
2020/02/26 14:07:59 Value: [a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4]  R=25d8c347fb71a839c1274ec2ebe666b8dd4a3255403d991be39bb10016aa51f1
2020/02/26 14:07:59 Value: [a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4] LR=a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4
2020/02/26 14:07:59 Value: [c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d]  R=a97c1a50c069d3397e9702f28ab38a420b537c6e942b8546860f8fb7a49c70e4
2020/02/26 14:07:59 Value: [c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d] LR=c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d
2020/02/26 14:07:59 Value: [10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5]  L=c54c5350dba05d216175318dbfa6dd732d1adbdc92beaf320443f523f822f01d
2020/02/26 14:07:59 Value: [5b2e60ccb6907fae79d249bf7943b050c28b5fef84b9b44399f58461dd120c4a]
2020/02/26 14:07:59 Value: [10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5]  R=5b2e60ccb6907fae79d249bf7943b050c28b5fef84b9b44399f58461dd120c4a
2020/02/26 14:07:59 Value: [10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5] LR=10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5
2020/02/26 14:07:59 Value: [4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e]  L=10555d6d9b28b8cc551cd2dd99fb270630b2cf2f56f8d70a7893a64b4b57ecb5
2020/02/26 14:07:59 Value: [c47e581dd90c10aad2cb4121c1a94a3199f9788f4f2e5d0f200f8f167ae2517c]
2020/02/26 14:07:59 Value: [4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e]  R=c47e581dd90c10aad2cb4121c1a94a3199f9788f4f2e5d0f200f8f167ae2517c
2020/02/26 14:07:59 Value: [4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e] LR=4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e
2020/02/26 14:07:59 Value: [b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f]  L=4b2b83117d905892053e5f4a4d28316a0107b9a360fce4633361e489b7b1359e
2020/02/26 14:07:59 Value: [dbdfc7af18a876d64d6fa2f0110e5dc6ead648f3e33f9e0163ceb267d47cd90d]
2020/02/26 14:07:59 Value: [b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f]  R=dbdfc7af18a876d64d6fa2f0110e5dc6ead648f3e33f9e0163ceb267d47cd90d
2020/02/26 14:07:59 Value: [b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f] LR=b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f
2020/02/26 14:07:59 Value: [9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808]  L=b918032b5cc7185ee2de8d38907e43b9ef9ffb61b5424ebf8a70643721c0af2f
2020/02/26 14:07:59 Value: [98f7dafd1f8c6ec84a555819d9cd05baed657e3973f86f3fb178dbd13c676bba]
2020/02/26 14:07:59 Value: [9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808]  R=98f7dafd1f8c6ec84a555819d9cd05baed657e3973f86f3fb178dbd13c676bba
2020/02/26 14:07:59 Value: [9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808] LR=9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808
2020/02/26 14:07:59 Value: [ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388]  L=9d7c106e64d5e6c40daacb4bcf4397912e31a78fe0e81b9fb4d79da08c7b2808
2020/02/26 14:07:59 Value: [36447318383768a406815764e128d404ca1482ec57cd995e12ac97962b33b73e]
2020/02/26 14:07:59 Value: [ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388]  R=36447318383768a406815764e128d404ca1482ec57cd995e12ac97962b33b73e
2020/02/26 14:07:59 Value: [ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388] LR=ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388
2020/02/26 14:07:59 Value: [e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e]  L=ee7536d1855cddf24dba36bfc04cdc1db522cef46e49c4d5aa5bde25648db388
2020/02/26 14:07:59 Value: [6e7ca4d2ea8919396dfd66a7d57338cb0ce8bca76043b5a4745e463c2c9d9af8]
2020/02/26 14:07:59 Value: [e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e]  R=6e7ca4d2ea8919396dfd66a7d57338cb0ce8bca76043b5a4745e463c2c9d9af8
2020/02/26 14:07:59 Value: [e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e] LR=e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e
2020/02/26 14:07:59 Value: [69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950]  L=e4d9506f6a877a6f15c4e240e53327cc7c38ef37e610ebfc62ce674308b3263e
2020/02/26 14:07:59 Value: [0f1301c53b6885f23037a81da41f450968a2b133f28b30d65d1fb431da74a9c1]
2020/02/26 14:07:59 Value: [69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950]  R=0f1301c53b6885f23037a81da41f450968a2b133f28b30d65d1fb431da74a9c1
2020/02/26 14:07:59 Value: [69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950] LR=69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950
2020/02/26 14:07:59 Correct: 69af555efcc31073fcb82977a5beb6fd84f20fe01f663cb1585c84945193f950
2020/02/26 14:07:59 1: 04D5CDB17FE99AB803F7D695CDAD7B4FC2D18D6C79FBBEFDF26115C142C5A3FBA961E9C00A7E448AD605B45021B98F2BC9509E930F51586B950B11C8B1A97A09FE
2020/02/26 14:07:59 2: 1299336a09d0452cfb5adbbde852e96dcc64c155f67fd2aa6568c448dfb1883c
2020/02/26 14:07:59 3: 89fd2d355c875644d14dc977eb41e9d4c0f32a39
2020/02/26 14:07:59 4: 0089fd2d355c875644d14dc977eb41e9d4c0f32a39
2020/02/26 14:07:59 5: 47468fe58b1728d0e7c9b0786d453bf8084b2d870c4bcea025916b6dd1610c7e
2020/02/26 14:07:59 6: 32a7454c269e5433f17fee57adb995e5639febc50e4b7ade855d0bd07b061a9c
2020/02/26 14:07:59 7: 32a7454c
2020/02/26 14:07:59 8: 0089fd2d355c875644d14dc977eb41e9d4c0f32a3932a7454c
2020/02/26 14:07:59 9: 1DacsJLsyr77YYkWpNGg1H5Mq3mUf84VzB
```