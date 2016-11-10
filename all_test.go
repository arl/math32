package math32

import (
	"math"
	"testing"
)

var vf = []float32{
	4.9790119248e+00,
	7.7388724745e+00,
	-2.7688005719e-01,
	-5.0106036189e+00,
	9.6362937071e+00,
	2.9263772392e+00,
	5.2290834314e+00,
	2.7279399104e+00,
	1.8253080916e+00,
	-8.6859247683e+00,
}

// The expected results below were computed by the high precision calculators
// at http://keisan.casio.com/.  More exact input values (array vf[], above)
// were obtained by printing them with "%.26f".  The answers were calculated
// to 26 digits (by using the "Digit number" drop-down control of each
// calculator).
var acos = []float32{
	1.0496193e+00,
	6.858401e-01,
	1.5984878e+00,
	2.0956199e+00,
	2.705301e-01,
	1.2738122e+00,
	1.0205369e+00,
	1.2945003e+00,
	1.3872365e+00,
	2.6231510e+00,
}
var acosh = []float32{
	2.47433470e+00,
	2.85763853e+00,
	7.27969615e-01,
	2.47967944e+00,
	3.05520207e+00,
	2.04423859e+00,
	2.51587015e+00,
	1.9905083e+00,
	1.698862579e+00,
	2.961145484e+00,
}
var asin = []float32{
	5.211769721e-01,
	8.8495624e-01,
	-2.7691545e-02,
	-5.24823609e-01,
	1.30026624e+00,
	2.96984158e-01,
	5.5025935e-01,
	2.7629596e-01,
	1.83559892e-01,
	-1.0523547e+00,
}
var asinh = []float32{
	2.30831391249e+00,
	2.74355159430 + 00,
	-2.73459085347e-01,
	-2.31451576445e+00,
	2.96136521540e+00,
	1.7949042e+00,
	2.35640329059e+00,
	1.72871187907e+00,
	1.36266580837e+00,
	-2.85814836264e+00,
}
var atan = []float32{
	1.37259026212e+00,
	1.44229060964e+00,
	-2.701132435995e-01,
	-1.373807768431e+00,
	1.467392119351e+00,
	1.241517356584e+00,
	1.381839686568e+00,
	1.219430584466e+00,
	1.069603195234e+00,
	-1.456172193879e+00,
}
var atanh = []float32{
	5.4651163712251e-01,
	1.0299475e+00,
	-2.769508442074e-02,
	-5.507209611920e-01,
	1.9943940993171e+00,
	3.0144860457808e-01,
	5.8033427206942e-01,
	2.7987996e-01,
	1.8459947964298e-01,
	-1.3273185e+00,
}
var atan2 = []float32{
	1.1088291730037004444527075e+00,
	9.1218183188715804018797795e-01,
	1.5984772603216203736068915e+00,
	2.0352918654092086637227327e+00,
	8.0391819139044720267356014e-01,
	1.2861075249894661588866752e+00,
	1.0889904479131695712182587e+00,
	1.3044821793397925293797357e+00,
	1.3902530903455392306872261e+00,
	2.2859857424479142655411058e+00,
}
var cbrt = []float32{
	1.7075799841925094446722675e+00,
	1.9779982212970353936691498e+00,
	-6.5177429017779910853339447e-01,
	-1.7111838886544019873338113e+00,
	2.1279920909827937423960472e+00,
	1.4303536770460741452312367e+00,
	1.7357021059106154902341052e+00,
	1.3972633462554328350552916e+00,
	1.2221149580905388454977636e+00,
	-2.0556003730500069110343596e+00,
}
var ceil = []float32{
	5.0000000000000000e+00,
	8.0000000000000000e+00,
	0.0000000000000000e+00,
	-5.0000000000000000e+00,
	1.0000000000000000e+01,
	3.0000000000000000e+00,
	6.0000000000000000e+00,
	3.0000000000000000e+00,
	2.0000000000000000e+00,
	-8.0000000000000000e+00,
}
var copysign = []float32{
	-4.9790119248836735e+00,
	-7.7388724745781045e+00,
	-2.7688005719200159e-01,
	-5.0106036182710749e+00,
	-9.6362937071984173e+00,
	-2.9263772392439646e+00,
	-5.2290834314593066e+00,
	-2.7279399104360102e+00,
	-1.8253080916808550e+00,
	-8.6859247685756013e+00,
}
var cos = []float32{
	2.634753e-01,
	1.1485507e-01,
	9.6191297325640768154550453e-01,
	2.9381394e-01,
	-9.777139e-01,
	-9.7693041344303219127199518e-01,
	4.940089e-01,
	-9.1565865e-01,
	-2.517729e-01,
	-7.392412e-01,
}

// Results for 100000 * Pi + vf[i]
var cosLarge = []float32{
	2.634752141185559426744e-01,
	1.14855126055543100712e-01,
	9.61912973266488928113e-01,
	2.9381411499556122552e-01,
	-9.777138189880161924641e-01,
	-9.76930413445147608049e-01,
	4.940088097314976789841e-01,
	-9.15658690217517835002e-01,
	-2.51772931436786954751e-01,
	-7.3924135157173099849e-01,
}
var cosh = []float32{
	7.26687969e+01,
	1.1479414e+03,
	1.03857679e+00,
	7.5000946e+01,
	7.6552446e+03,
	9.356749175e+00,
	9.331352e+01,
	7.6833425e+00,
	3.182937162e+00,
	2.9595051e+03,
}
var erf = []float32{
	5.18653548177e-01,
	7.26238758341e-01,
	-3.1234586882e-02,
	-5.2143121110e-01,
	8.27047426713e-01,
	3.21017675583e-01,
	5.40399031222e-01,
	3.00347029167e-01,
	2.03699244178e-01,
	-7.806938e-01,
}
var erfc = []float32{
	4.81346451e-01,
	2.737612e-01,
	1.03123458e+00,
	1.52143121e+00,
	1.72952573e-01,
	6.78982324e-01,
	4.5960099e-01,
	6.99652970e-01,
	7.96300755e-01,
	1.78069386e+00,
}
var exp = []float32{
	1.45330713e+02,
	2.29588225e+03,
	7.58145425e-01,
	6.666879e-03,
	1.5310488e+04,
	1.86599075e+01,
	1.8662169e+02,
	1.53013315e+01,
	6.20470634e+00,
	1.6894717e-04,
}
var expm1 = []float32{
	5.105048e-02,
	8.0462e-02,
	-2.764970978891639815187418703e-03,
	-4.887143e-02,
	1.0115864277221467777117227494e-01,
	2.969616407795910726014621657e-02,
	5.368214487944892300914037972e-02,
	2.765488851131274068067445335e-02,
	1.842068661871398836913874273e-02,
	-8.3193870863553801814961137573e-02,
}
var expm1Large = []float32{
	4.2031414e+21,
	4.069081e+33,
	-0.9372627915981363e+00,
	-1.0,
	Inf(1),
	5.117939e+12,
	5.124148e+22,
	7.035452e+11,
	8.456917e+07,
	-1.0,
}
var exp2 = []float32{
	3.1537842e+01,
	2.13615492e+02,
	8.25374025e-01,
	3.1021163e-02,
	7.9581726e+02,
	7.60199058e+00,
	3.7506886e+01,
	6.62508934e+00,
	3.54382679e+00,
	2.4281538e-03,
}
var fabs = []float32{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	2.7688005719200159e-01,
	5.0106036182710749e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	8.6859247685756013e+00,
}
var fdim = []float32{
	4.9790119248836735e+00,
	7.7388724745781045e+00,
	0.0000000000000000e+00,
	0.0000000000000000e+00,
	9.6362937071984173e+00,
	2.9263772392439646e+00,
	5.2290834314593066e+00,
	2.7279399104360102e+00,
	1.8253080916808550e+00,
	0.0000000000000000e+00,
}
var floor = []float32{
	4.0000000000000000e+00,
	7.0000000000000000e+00,
	-1.0000000000000000e+00,
	-6.0000000000000000e+00,
	9.0000000000000000e+00,
	2.0000000000000000e+00,
	5.0000000000000000e+00,
	2.0000000000000000e+00,
	1.0000000000000000e+00,
	-9.0000000000000000e+00,
}
var fmod = []float32{
	4.1975975e-02,
	2.2611275e+00,
	3.2317996e-02,
	4.9893966e+00,
	3.637066e-01,
	1.2208681e+00,
	4.7709165e+00,
	1.8161805e+00,
	8.734596e-01,
	1.3140755e+00,
}

type fi struct {
	f float32
	i int
}

var frexp = []fi{
	{6.2237649061045918750e-01, 3},
	{9.6735905932226306250e-01, 3},
	{-5.5376011438400318000e-01, -1},
	{-6.2632545228388436250e-01, 3},
	{6.02268356699901081250e-01, 4},
	{7.3159430981099115000e-01, 2},
	{6.5363542893241332500e-01, 3},
	{6.8198497760900255000e-01, 2},
	{9.1265404584042750000e-01, 1},
	{-5.4287029803597508250e-01, 4},
}
var gamma = []float32{
	2.325435e+01,
	2.991154e+03,
	-4.5611543e+00,
	7.719545e-01,
	1.6111866e+05,
	1.8706576e+00,
	3.4082794e+01,
	1.5797338e+00,
	9.38345869e-01,
	-2.093996e-05,
}
var j0 = []float32{
	-1.8444682230601672018219338e-01,
	2.27353668906331975435892e-01,
	9.809259936157051116270273e-01,
	-1.741170131426226587841181e-01,
	-2.1389448451144143352039069e-01,
	-2.340905848928038763337414e-01,
	-1.0029099691890912094586326e-01,
	-1.5466726714884328135358907e-01,
	3.252650187653420388714693e-01,
	-8.72218484409407250005360235e-03,
}
var j1 = []float32{
	-3.251526395295203422162967e-01,
	1.893581711430515718062564e-01,
	-1.3711761352467242914491514e-01,
	3.287486536269617297529617e-01,
	1.3133899188830978473849215e-01,
	3.660243417832986825301766e-01,
	-3.4436769271848174665420672e-01,
	4.329481396640773768835036e-01,
	5.8181350531954794639333955e-01,
	-2.7030574577733036112996607e-01,
}
var j2 = []float32{
	5.3837518920137802565192769e-02,
	-1.7841678003393207281244667e-01,
	9.521746934916464142495821e-03,
	4.28958355470987397983072e-02,
	2.4115371837854494725492872e-01,
	4.842458532394520316844449e-01,
	-3.142145220618633390125946e-02,
	4.720849184745124761189957e-01,
	3.122312022520957042957497e-01,
	7.096213118930231185707277e-02,
}
var jM3 = []float32{
	-3.684042080996403091021151e-01,
	2.8157665936340887268092661e-01,
	4.401005480841948348343589e-04,
	3.629926999056814081597135e-01,
	3.123672198825455192489266e-02,
	-2.958805510589623607540455e-01,
	-3.2033177696533233403289416e-01,
	-2.592737332129663376736604e-01,
	-1.0241334641061485092351251e-01,
	-2.3762660886100206491674503e-01,
}
var lgamma = []fi{
	{3.146492141244545774319734e+00, 1},
	{8.003414490659126375852113e+00, 1},
	{1.517575735509779707488106e+00, -1},
	{-2.588480028182145853558748e-01, 1},
	{1.1989897050205555002007985e+01, 1},
	{6.262899811091257519386906e-01, 1},
	{3.5287924899091566764846037e+00, 1},
	{4.5725644770161182299423372e-01, 1},
	{-6.363667087767961257654854e-02, 1},
	{-1.077385130910300066425564e+01, -1},
}
var log = []float32{
	1.6052315e+00,
	2.0462560018708770653153909e+00,
	-1.2841708730962657801275038e+00,
	1.6115563905281545116286206e+00,
	2.2655365644872016636317461e+00,
	1.0737652208918379856272735e+00,
	1.6542361 + 00,
	1.0035467127723465801264487e+00,
	6.0174879014578057187016475e-01,
	2.161703872847352815363655e+00,
}
var logb = []float32{
	2.0000000000000000e+00,
	2.0000000000000000e+00,
	-2.0000000000000000e+00,
	2.0000000000000000e+00,
	3.0000000000000000e+00,
	1.0000000000000000e+00,
	2.0000000000000000e+00,
	1.0000000000000000e+00,
	0.0000000000000000e+00,
	3.0000000000000000e+00,
}
var log10 = []float32{
	6.9714316642508290997617083e-01,
	8.886776901739320576279124e-01,
	-5.5770832400658929815908236e-01,
	6.998900476822994346229723e-01,
	9.8391002850684232013281033e-01,
	4.6633031029295153334285302e-01,
	7.1842557117242328821552533e-01,
	4.3583479968917773161304553e-01,
	2.6133617905227038228626834e-01,
	9.3881606348649405716214241e-01,
}
var log1p = []float32{
	4.8590257759797794104158205e-02,
	7.4540265965225865330849141e-02,
	-2.7726407903942672823234024e-03,
	-5.1404917651627649094953380e-02,
	9.1998280672258624681335010e-02,
	2.8843762576593352865894824e-02,
	5.0969534581863707268992645e-02,
	2.6913947602193238458458594e-02,
	1.8088493239630770262045333e-02,
	-9.0865245631588989681559268e-02,
}
var log2 = []float32{
	2.3158594707062190618898251e+00,
	2.9521233862883917703341018e+00,
	-1.8526669502700329984917062e+00,
	2.3249844127278861543568029e+00,
	3.268478366538305087466309e+00,
	1.5491157592596970278166492e+00,
	2.3865580889631732407886495e+00,
	1.447811865817085365540347e+00,
	8.6813999540425116282815557e-01,
	3.118679457227342224364709e+00,
}
var modf = [][2]float32{
	{4.0000000000000000e+00, 9.7901192488367350108546816e-01},
	{7.0000000000000000e+00, 7.3887247457810456552351752e-01},
	{Copysign(0, -1), -2.7688005719200159404635997e-01},
	{-5.0000000000000000e+00, -1.060361827107492160848778e-02},
	{9.0000000000000000e+00, 6.3629370719841737980004837e-01},
	{2.0000000000000000e+00, 9.2637723924396464525443662e-01},
	{5.0000000000000000e+00, 2.2908343145930665230025625e-01},
	{2.0000000000000000e+00, 7.2793991043601025126008608e-01},
	{1.0000000000000000e+00, 8.2530809168085506044576505e-01},
	{-8.0000000000000000e+00, -6.8592476857560136238589621e-01},
}
var nextafter32 = []float32{
	4.979012489318848e+00,
	7.738873004913330e+00,
	-2.768800258636475e-01,
	-5.010602951049805e+00,
	9.636294364929199e+00,
	2.926377534866333e+00,
	5.229084014892578e+00,
	2.727940082550049e+00,
	1.825308203697205e+00,
	-8.685923576354980e+00,
}
var nextafter64 = []float32{
	4.97901192488367438926388786e+00,
	7.73887247457810545370193722e+00,
	-2.7688005719200153853520874e-01,
	-5.01060361827107403343006808e+00,
	9.63629370719841915615688777e+00,
	2.92637723924396508934364647e+00,
	5.22908343145930754047867595e+00,
	2.72793991043601069534929593e+00,
	1.82530809168085528249036997e+00,
	-8.68592476857559958602905681e+00,
}
var pow = []float32{
	9.528225e+04,
	5.481161e+07,
	5.2859121715894396531132279e-01,
	9.758804e-06,
	4.3280614e+09,
	8.4406775e+02,
	1.6946638e+05,
	5.344903e+02,
	6.688182138451414936380374e+01,
	2.060988e-09,
}
var remainder = []float32{
	4.197615023265299782906368e-02,
	2.261127525421895434476482e+00,
	3.231794108794261433104108e-02,
	-2.120723654214984321697556e-02,
	3.637062928015826201999516e-01,
	1.220868282268106064236690e+00,
	-4.581668629186133046005125e-01,
	-9.117596417440410050403443e-01,
	8.734595415957246977711748e-01,
	1.314075231424398637614104e+00,
}
var signbit = []bool{
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
}
var sin = []float32{
	-9.6466616586009283766724726e-01,
	9.9338225271646545763467022e-01,
	-2.7335587039794393342449301e-01,
	9.5586257685042792878173752e-01,
	-2.099421066779969164496634e-01,
	2.135578780799860532750616e-01,
	-8.694568971167362743327708e-01,
	4.019566681155577786649878e-01,
	9.6778633541687993721617774e-01,
	-6.734405869050344734943028e-01,
}

// Results for 100000 * Pi + vf[i]
var sinLarge = []float32{
	-9.646661658548936063912e-01,
	9.933822527198506903752e-01,
	-2.7335587036246899796e-01,
	9.55862576853689321268e-01,
	-2.099421066862688873691e-01,
	2.13557878070308981163e-01,
	-8.694568970959221300497e-01,
	4.01956668098863248917e-01,
	9.67786335404528727927e-01,
	-6.7344058693131973066e-01,
}
var sinh = []float32{
	7.2661916084208532301448439e+01,
	1.1479409110035194500526446e+03,
	-2.8043136512812518927312641e-01,
	-7.499429091181587232835164e+01,
	7.6552466042906758523925934e+03,
	9.3031583421672014313789064e+00,
	9.330815755828109072810322e+01,
	7.6179893137269146407361477e+00,
	3.021769180549615819524392e+00,
	-2.95950575724449499189888e+03,
}
var sqrt = []float32{
	2.2313699659365484748756904e+00,
	2.7818829009464263511285458e+00,
	5.2619393496314796848143251e-01,
	2.2384377628763938724244104e+00,
	3.1042380236055381099288487e+00,
	1.7106657298385224403917771e+00,
	2.286718922705479046148059e+00,
	1.6516476e+00,
	1.3510396336454586262419247e+00,
	2.9471892997524949215723329e+00,
}
var tan = []float32{
	-3.661316565040227801781974e+00,
	8.64900232648597589369854e+00,
	-2.8417941955033612725238097e-01,
	3.253290185974728640827156e+00,
	2.147275640380293804770778e-01,
	-2.18600910711067004921551e-01,
	-1.760002817872367935518928e+00,
	-4.389808914752818126249079e-01,
	-3.843885560201130679995041e+00,
	9.10988793377685105753416e-01,
}

// Results for 100000 * Pi + vf[i]
var tanLarge = []float32{
	-3.66131656475596512705e+00,
	8.6490023287202547927e+00,
	-2.841794195104782406e-01,
	3.2532901861033120983e+00,
	2.14727564046880001365e-01,
	-2.18600910700688062874e-01,
	-1.760002817699722747043e+00,
	-4.38980891453536115952e-01,
	-3.84388555942723509071e+00,
	9.1098879344275101051e-01,
}
var tanh = []float32{
	9.9990531206936338549262119e-01,
	9.9999962057085294197613294e-01,
	-2.7001505097318677233756845e-01,
	-9.9991110943061718603541401e-01,
	9.9999999146798465745022007e-01,
	9.9427249436125236705001048e-01,
	9.9994257600983138572705076e-01,
	9.9149409509772875982054701e-01,
	9.4936501296239685514466577e-01,
	-9.9999994291374030946055701e-01,
}
var trunc = []float32{
	4.0000000000000000e+00,
	7.0000000000000000e+00,
	-0.0000000000000000e+00,
	-5.0000000000000000e+00,
	9.0000000000000000e+00,
	2.0000000000000000e+00,
	5.0000000000000000e+00,
	2.0000000000000000e+00,
	1.0000000000000000e+00,
	-8.0000000000000000e+00,
}
var y0 = []float32{
	-3.053399153780788357534855e-01,
	1.7437227649515231515503649e-01,
	-8.6221781263678836910392572e-01,
	-3.100664880987498407872839e-01,
	1.422200649300982280645377e-01,
	4.000004067997901144239363e-01,
	-3.3340749753099352392332536e-01,
	4.5399790746668954555205502e-01,
	4.8290004112497761007536522e-01,
	2.7036697826604756229601611e-01,
}
var y1 = []float32{
	0.15494213737457922210218611,
	-0.2165955142081145245075746,
	-2.4644949631241895201032829,
	0.1442740489541836405154505,
	0.2215379960518984777080163,
	0.3038800915160754150565448,
	0.0691107642452362383808547,
	0.2380116417809914424860165,
	-0.20849492979459761009678934,
	0.0242503179793232308250804,
}
var y2 = []float32{
	0.3675780219390303613394936,
	-0.23034826393250119879267257,
	-16.939677983817727205631397,
	0.367653980523052152867791,
	-0.0962401471767804440353136,
	-0.1923169356184851105200523,
	0.35984072054267882391843766,
	-0.2794987252299739821654982,
	-0.7113490692587462579757954,
	-0.2647831587821263302087457,
}
var yM3 = []float32{
	-0.14035984421094849100895341,
	-0.097535139617792072703973,
	242.25775994555580176377379,
	-0.1492267014802818619511046,
	0.26148702629155918694500469,
	0.56675383593895176530394248,
	-0.206150264009006981070575,
	0.64784284687568332737963658,
	1.3503631555901938037008443,
	0.1461869756579956803341844,
}

// arguments and expected results for special cases
var vfacosSC = []float32{
	-Pi,
	1,
	Pi,
	NaN(),
}
var acosSC = []float32{
	NaN(),
	0,
	NaN(),
	NaN(),
}

var vfacoshSC = []float32{
	Inf(-1),
	0.5,
	1,
	Inf(1),
	NaN(),
}
var acoshSC = []float32{
	NaN(),
	NaN(),
	0,
	Inf(1),
	NaN(),
}

var vfasinSC = []float32{
	-Pi,
	Copysign(0, -1),
	0,
	Pi,
	NaN(),
}
var asinSC = []float32{
	NaN(),
	Copysign(0, -1),
	0,
	NaN(),
	NaN(),
}

var vfasinhSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var asinhSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}

var vfatanSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var atanSC = []float32{
	-Pi / 2,
	Copysign(0, -1),
	0,
	Pi / 2,
	NaN(),
}

var vfatanhSC = []float32{
	Inf(-1),
	-Pi,
	-1,
	Copysign(0, -1),
	0,
	1,
	Pi,
	Inf(1),
	NaN(),
}
var atanhSC = []float32{
	NaN(),
	NaN(),
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
	NaN(),
	NaN(),
}
var vfatan2SC = [][2]float32{
	{Inf(-1), Inf(-1)},
	{Inf(-1), -Pi},
	{Inf(-1), 0},
	{Inf(-1), +Pi},
	{Inf(-1), Inf(1)},
	{Inf(-1), NaN()},
	{-Pi, Inf(-1)},
	{-Pi, 0},
	{-Pi, Inf(1)},
	{-Pi, NaN()},
	{Copysign(0, -1), Inf(-1)},
	{Copysign(0, -1), -Pi},
	{Copysign(0, -1), Copysign(0, -1)},
	{Copysign(0, -1), 0},
	{Copysign(0, -1), +Pi},
	{Copysign(0, -1), Inf(1)},
	{Copysign(0, -1), NaN()},
	{0, Inf(-1)},
	{0, -Pi},
	{0, Copysign(0, -1)},
	{0, 0},
	{0, +Pi},
	{0, Inf(1)},
	{0, NaN()},
	{+Pi, Inf(-1)},
	{+Pi, 0},
	{+Pi, Inf(1)},
	{+Pi, NaN()},
	{Inf(1), Inf(-1)},
	{Inf(1), -Pi},
	{Inf(1), 0},
	{Inf(1), +Pi},
	{Inf(1), Inf(1)},
	{Inf(1), NaN()},
	{NaN(), NaN()},
}
var atan2SC = []float32{
	-3 * Pi / 4,     // atan2(-Inf, -Inf)
	-Pi / 2,         // atan2(-Inf, -Pi)
	-Pi / 2,         // atan2(-Inf, +0)
	-Pi / 2,         // atan2(-Inf, +Pi)
	-Pi / 4,         // atan2(-Inf, +Inf)
	NaN(),           // atan2(-Inf, NaN)
	-Pi,             // atan2(-Pi, -Inf)
	-Pi / 2,         // atan2(-Pi, +0)
	Copysign(0, -1), // atan2(-Pi, Inf)
	NaN(),           // atan2(-Pi, NaN)
	-Pi,             // atan2(-0, -Inf)
	-Pi,             // atan2(-0, -Pi)
	-Pi,             // atan2(-0, -0)
	Copysign(0, -1), // atan2(-0, +0)
	Copysign(0, -1), // atan2(-0, +Pi)
	Copysign(0, -1), // atan2(-0, +Inf)
	NaN(),           // atan2(-0, NaN)
	Pi,              // atan2(+0, -Inf)
	Pi,              // atan2(+0, -Pi)
	Pi,              // atan2(+0, -0)
	0,               // atan2(+0, +0)
	0,               // atan2(+0, +Pi)
	0,               // atan2(+0, +Inf)
	NaN(),           // atan2(+0, NaN)
	Pi,              // atan2(+Pi, -Inf)
	Pi / 2,          // atan2(+Pi, +0)
	0,               // atan2(+Pi, +Inf)
	NaN(),           // atan2(+Pi, NaN)
	3 * Pi / 4,      // atan2(+Inf, -Inf)
	Pi / 2,          // atan2(+Inf, -Pi)
	Pi / 2,          // atan2(+Inf, +0)
	Pi / 2,          // atan2(+Inf, +Pi)
	Pi / 4,          // atan2(+Inf, +Inf)
	NaN(),           // atan2(+Inf, NaN)
	NaN(),           // atan2(NaN, NaN)
}

var vfcbrtSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var cbrtSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}

var vfceilSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var ceilSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}

var vfcopysignSC = []float32{
	Inf(-1),
	Inf(1),
	NaN(),
}
var copysignSC = []float32{
	Inf(-1),
	Inf(-1),
	NaN(),
}

var vfcosSC = []float32{
	Inf(-1),
	Inf(1),
	NaN(),
}
var cosSC = []float32{
	NaN(),
	NaN(),
	NaN(),
}

var vfcoshSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var coshSC = []float32{
	Inf(1),
	1,
	1,
	Inf(1),
	NaN(),
}

var vferfSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var erfSC = []float32{
	-1,
	Copysign(0, -1),
	0,
	1,
	NaN(),
}

var vferfcSC = []float32{
	Inf(-1),
	Inf(1),
	NaN(),
}
var erfcSC = []float32{
	2,
	0,
	NaN(),
}

var vfexpSC = []float32{
	Inf(-1),
	-2000,
	2000,
	Inf(1),
	NaN(),
}
var expSC = []float32{
	0,
	0,
	Inf(1),
	Inf(1),
	NaN(),
}

var vfexpm1SC = []float32{
	Inf(-1),
	-710,
	Copysign(0, -1),
	0,
	710,
	Inf(1),
	NaN(),
}
var expm1SC = []float32{
	-1,
	-1,
	Copysign(0, -1),
	0,
	Inf(1),
	Inf(1),
	NaN(),
}

var vffabsSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var fabsSC = []float32{
	Inf(1),
	0,
	0,
	Inf(1),
	NaN(),
}

var vffdimSC = [][2]float32{
	{Inf(-1), Inf(-1)},
	{Inf(-1), Inf(1)},
	{Inf(-1), NaN()},
	{Copysign(0, -1), Copysign(0, -1)},
	{Copysign(0, -1), 0},
	{0, Copysign(0, -1)},
	{0, 0},
	{Inf(1), Inf(-1)},
	{Inf(1), Inf(1)},
	{Inf(1), NaN()},
	{NaN(), Inf(-1)},
	{NaN(), Copysign(0, -1)},
	{NaN(), 0},
	{NaN(), Inf(1)},
	{NaN(), NaN()},
}

//var nan = float32frombits(0xFFF8000000000000) // SSE2 DIVSD 0/0
var nan = NaN()
var vffdim2SC = [][2]float32{
	{Inf(-1), Inf(-1)},
	{Inf(-1), Inf(1)},
	{Inf(-1), nan},
	{Copysign(0, -1), Copysign(0, -1)},
	{Copysign(0, -1), 0},
	{0, Copysign(0, -1)},
	{0, 0},
	{Inf(1), Inf(-1)},
	{Inf(1), Inf(1)},
	{Inf(1), nan},
	{nan, Inf(-1)},
	{nan, Copysign(0, -1)},
	{nan, 0},
	{nan, Inf(1)},
	{nan, nan},
}
var fdimSC = []float32{
	NaN(),
	0,
	NaN(),
	0,
	0,
	0,
	0,
	Inf(1),
	NaN(),
	NaN(),
	NaN(),
	NaN(),
	NaN(),
	NaN(),
	NaN(),
}
var fmaxSC = []float32{
	Inf(-1),
	Inf(1),
	NaN(),
	Copysign(0, -1),
	0,
	0,
	0,
	Inf(1),
	Inf(1),
	Inf(1),
	NaN(),
	NaN(),
	NaN(),
	Inf(1),
	NaN(),
}
var fminSC = []float32{
	Inf(-1),
	Inf(-1),
	Inf(-1),
	Copysign(0, -1),
	Copysign(0, -1),
	Copysign(0, -1),
	0,
	Inf(-1),
	Inf(1),
	NaN(),
	Inf(-1),
	NaN(),
	NaN(),
	NaN(),
	NaN(),
}

var vffmodSC = [][2]float32{
	{Inf(-1), Inf(-1)},
	{Inf(-1), -Pi},
	{Inf(-1), 0},
	{Inf(-1), Pi},
	{Inf(-1), Inf(1)},
	{Inf(-1), NaN()},
	{-Pi, Inf(-1)},
	{-Pi, 0},
	{-Pi, Inf(1)},
	{-Pi, NaN()},
	{Copysign(0, -1), Inf(-1)},
	{Copysign(0, -1), 0},
	{Copysign(0, -1), Inf(1)},
	{Copysign(0, -1), NaN()},
	{0, Inf(-1)},
	{0, 0},
	{0, Inf(1)},
	{0, NaN()},
	{Pi, Inf(-1)},
	{Pi, 0},
	{Pi, Inf(1)},
	{Pi, NaN()},
	{Inf(1), Inf(-1)},
	{Inf(1), -Pi},
	{Inf(1), 0},
	{Inf(1), Pi},
	{Inf(1), Inf(1)},
	{Inf(1), NaN()},
	{NaN(), Inf(-1)},
	{NaN(), -Pi},
	{NaN(), 0},
	{NaN(), Pi},
	{NaN(), Inf(1)},
	{NaN(), NaN()},
}
var fmodSC = []float32{
	NaN(),           // fmod(-Inf, -Inf)
	NaN(),           // fmod(-Inf, -Pi)
	NaN(),           // fmod(-Inf, 0)
	NaN(),           // fmod(-Inf, Pi)
	NaN(),           // fmod(-Inf, +Inf)
	NaN(),           // fmod(-Inf, NaN)
	-Pi,             // fmod(-Pi, -Inf)
	NaN(),           // fmod(-Pi, 0)
	-Pi,             // fmod(-Pi, +Inf)
	NaN(),           // fmod(-Pi, NaN)
	Copysign(0, -1), // fmod(-0, -Inf)
	NaN(),           // fmod(-0, 0)
	Copysign(0, -1), // fmod(-0, Inf)
	NaN(),           // fmod(-0, NaN)
	0,               // fmod(0, -Inf)
	NaN(),           // fmod(0, 0)
	0,               // fmod(0, +Inf)
	NaN(),           // fmod(0, NaN)
	Pi,              // fmod(Pi, -Inf)
	NaN(),           // fmod(Pi, 0)
	Pi,              // fmod(Pi, +Inf)
	NaN(),           // fmod(Pi, NaN)
	NaN(),           // fmod(+Inf, -Inf)
	NaN(),           // fmod(+Inf, -Pi)
	NaN(),           // fmod(+Inf, 0)
	NaN(),           // fmod(+Inf, Pi)
	NaN(),           // fmod(+Inf, +Inf)
	NaN(),           // fmod(+Inf, NaN)
	NaN(),           // fmod(NaN, -Inf)
	NaN(),           // fmod(NaN, -Pi)
	NaN(),           // fmod(NaN, 0)
	NaN(),           // fmod(NaN, Pi)
	NaN(),           // fmod(NaN, +Inf)
	NaN(),           // fmod(NaN, NaN)
}

var vffrexpSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var frexpSC = []fi{
	{Inf(-1), 0},
	{Copysign(0, -1), 0},
	{0, 0},
	{Inf(1), 0},
	{NaN(), 0},
}

var vfgamma = [][2]float32{
	{Inf(1), Inf(1)},
	{Inf(-1), NaN()},
	{0, Inf(1)},
	{Copysign(0, -1), Inf(-1)},
	{NaN(), NaN()},
	{-1, NaN()},
	{-2, NaN()},
	{-3, NaN()},
	{-1e16, NaN()},
	//{-1e300, NaN()},
	//{1.7e308, Inf(1)},

	// Test inputs inspired by Python test suite.
	// Outputs computed at high precision by PARI/GP.
	// If recomputing table entries, be careful to use
	// high-precision (%.1000g) formatting of the float64 inputs.
	// For example, -2.0000000000000004 is the float64 with exact value
	// -2.00000000000000044408920985626161695, and
	// gamma(-2.0000000000000004) = -1249999999999999.5386078562728167651513, while
	// gamma(-2.00000000000000044408920985626161695) = -1125899906826907.2044875028130093136826.
	// Thus the table lists -1.1258999068426235e+15 as the answer.
	{0.5, 1.772453850905516},
	{1.5, 0.886226925452758},
	{2.5, 1.329340388179137},
	{3.5, 3.3233509704478426},
	{-0.5, -3.544907701811032},
	{-1.5, 2.363271801207355},
	{-2.5, -0.9453087204829419},
	{-3.5, 0.2700882058522691},
	{0.1, 9.51350769866873},
	{0.01, 99.4325851191506},
	{1e-08, 9.999999942278434e+07},
	{1e-16, 1e+16},
	{0.001, 999.4237},
	{1e-16, 1e+16},
	//{1e-308, 1e+308},
	//{5.6e-309, 1.7857142857142864e+308},
	{5.5e-309, Inf(1)},
	{1e-309, Inf(1)},
	{1e-323, Inf(1)},
	{5e-324, Inf(1)},
	{-0.1, -10.686287021193193},
	{-0.01, -100.587204},
	{-1e-08, -1.0000000057721567e+08},
	{-1e-16, -1e+16},
	{-0.001, -1000.5782056293586},
	{-1e-16, -1e+16},
	//{-1e-308, -1e+308},
	//{-5.6e-309, -1.7857142857142864e+308},
	{0, Inf(1)},
	{0, Inf(1)},
	{0, Inf(1)},
	{0, Inf(1)},
	{-0.9999999999999999, NaN()},
	{-1.0000000000000002, NaN()},
	{-1.9999999999999998, NaN()},
	{-2.0000000000000004, NaN()},
	{-100.00000000000001, NaN()},
	{-99.99999999999999, NaN()},
	{17, 2.0922789888e+13},
	//{171, 7.257415615307999e+306},
	//{171.6, 1.5858969096672565e+308},
	//{171.624, 1.7942117599248104e+308},
	{171.625, Inf(1)},
	{172, Inf(1)},
	{2000, Inf(1)},
	{-100.5, Copysign(0, -1)},
	{-160.5, Copysign(0, -1)},
	// those tests are commented because it need to be investigated by they
	// fail on travis:
	// go version devel +2925427 Tue Nov 8 23:06:17 2016 +0000 linux/amd64
	//{-170.5, -3.3127395215386074e-308},
	//{-171.5, Inf(1)},
	//{-176.5, Inf(1)},
	//{-177.5, Inf(1)},
	//{-178.5, Inf(1)},
	//{-179.5, Inf(1)},
	//{-201.0001, Inf(1)},
	//{-202.9999, Inf(1)},
	//{-1000.5, Inf(1)},
	{-1.0000000003e+09, NaN()},
	{-4.5035996273704955e+15, NaN()},
	{-63.349078729022985, 4.177797167776188e-88},
	{-127.45117632943295, 1.183111089623681e-214},
}

var vfhypotSC = [][2]float32{
	{Inf(-1), Inf(-1)},
	{Inf(-1), 0},
	{Inf(-1), Inf(1)},
	{Inf(-1), NaN()},
	{Copysign(0, -1), Copysign(0, -1)},
	{Copysign(0, -1), 0},
	{0, Copysign(0, -1)},
	{0, 0}, // +0, +0
	{0, Inf(-1)},
	{0, Inf(1)},
	{0, NaN()},
	{Inf(1), Inf(-1)},
	{Inf(1), 0},
	{Inf(1), Inf(1)},
	{Inf(1), NaN()},
	{NaN(), Inf(-1)},
	{NaN(), 0},
	{NaN(), Inf(1)},
	{NaN(), NaN()},
}
var hypotSC = []float32{
	Inf(1),
	Inf(1),
	Inf(1),
	Inf(1),
	0,
	0,
	0,
	0,
	Inf(1),
	Inf(1),
	NaN(),
	Inf(1),
	Inf(1),
	Inf(1),
	Inf(1),
	Inf(1),
	NaN(),
	Inf(1),
	NaN(),
}

var ilogbSC = []int{
	math.MaxInt32,
	math.MinInt32,
	math.MaxInt32,
	math.MaxInt32,
}

var vfj0SC = []float32{
	Inf(-1),
	0,
	Inf(1),
	NaN(),
}
var j0SC = []float32{
	0,
	1,
	0,
	NaN(),
}
var j1SC = []float32{
	0,
	0,
	0,
	NaN(),
}
var j2SC = []float32{
	0,
	0,
	0,
	NaN(),
}
var jM3SC = []float32{
	0,
	0,
	0,
	NaN(),
}

var vfldexpSC = []fi{
	{0, 0},
	{0, -1075},
	{0, 1024},
	{Copysign(0, -1), 0},
	{Copysign(0, -1), -1075},
	{Copysign(0, -1), 1024},
	{Inf(1), 0},
	{Inf(1), -1024},
	{Inf(-1), 0},
	{Inf(-1), -1024},
	{NaN(), -1024},
}
var ldexpSC = []float32{
	0,
	0,
	0,
	Copysign(0, -1),
	Copysign(0, -1),
	Copysign(0, -1),
	Inf(1),
	Inf(1),
	Inf(-1),
	Inf(-1),
	NaN(),
}

var vflgammaSC = []float32{
	Inf(-1),
	-3,
	0,
	1,
	2,
	Inf(1),
	NaN(),
}
var lgammaSC = []fi{
	{Inf(-1), 1},
	{Inf(1), 1},
	{Inf(1), 1},
	{0, 1},
	{0, 1},
	{Inf(1), 1},
	{NaN(), 1},
}

var vflogSC = []float32{
	Inf(-1),
	-Pi,
	Copysign(0, -1),
	0,
	1,
	Inf(1),
	NaN(),
}
var logSC = []float32{
	NaN(),
	NaN(),
	Inf(-1),
	Inf(-1),
	0,
	Inf(1),
	NaN(),
}

var vflogbSC = []float32{
	Inf(-1),
	0,
	Inf(1),
	NaN(),
}
var logbSC = []float32{
	Inf(1),
	Inf(-1),
	Inf(1),
	NaN(),
}

var vflog1pSC = []float32{
	Inf(-1),
	-Pi,
	-1,
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var log1pSC = []float32{
	NaN(),
	NaN(),
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}

var vfmodfSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	Inf(1),
	NaN(),
}
var modfSC = [][2]float32{
	{Inf(-1), NaN()}, // [2]float64{Copysign(0, -1), Inf(-1)},
	{Copysign(0, -1), Copysign(0, -1)},
	{Inf(1), NaN()}, // [2]float64{0, Inf(1)},
	{NaN(), NaN()},
}

var vfnextafter32SC = [][2]float32{
	{0, 0},
	{0, float32(Copysign(0, -1))},
	{0, -1},
	{0, float32(NaN())},
	{float32(Copysign(0, -1)), 1},
	{float32(Copysign(0, -1)), 0},
	{float32(Copysign(0, -1)), float32(Copysign(0, -1))},
	{float32(Copysign(0, -1)), -1},
	{float32(NaN()), 0},
	{float32(NaN()), float32(NaN())},
}
var nextafter32SC = []float32{
	0,
	0,
	-1.401298464e-45, // Float32frombits(0x80000001)
	float32(NaN()),
	1.401298464e-45, // Float32frombits(0x00000001)
	float32(Copysign(0, -1)),
	float32(Copysign(0, -1)),
	-1.401298464e-45, // Float32frombits(0x80000001)
	float32(NaN()),
	float32(NaN()),
}

var vfnextafter64SC = [][2]float32{
	{0, 0},
	{0, Copysign(0, -1)},
	{0, -1},
	{0, NaN()},
	{Copysign(0, -1), 1},
	{Copysign(0, -1), 0},
	{Copysign(0, -1), Copysign(0, -1)},
	{Copysign(0, -1), -1},
	{NaN(), 0},
	{NaN(), NaN()},
}
var nextafter64SC = []float32{
	0,
	0,
	-4.9406564584124654418e-324, // Float64frombits(0x8000000000000001)
	NaN(),
	4.9406564584124654418e-324, // Float64frombits(0x0000000000000001)
	Copysign(0, -1),
	Copysign(0, -1),
	-4.9406564584124654418e-324, // Float64frombits(0x8000000000000001)
	NaN(),
	NaN(),
}

var vfpowSC = [][2]float32{
	{Inf(-1), -Pi},
	{Inf(-1), -3},
	{Inf(-1), Copysign(0, -1)},
	{Inf(-1), 0},
	{Inf(-1), 1},
	{Inf(-1), 3},
	{Inf(-1), Pi},
	{Inf(-1), NaN()},

	{-Pi, Inf(-1)},
	{-Pi, -Pi},
	{-Pi, Copysign(0, -1)},
	{-Pi, 0},
	{-Pi, 1},
	{-Pi, Pi},
	{-Pi, Inf(1)},
	{-Pi, NaN()},

	{-1, Inf(-1)},
	{-1, Inf(1)},
	{-1, NaN()},
	{-1 / 2, Inf(-1)},
	{-1 / 2, Inf(1)},
	{Copysign(0, -1), Inf(-1)},
	{Copysign(0, -1), -Pi},
	{Copysign(0, -1), -3},
	{Copysign(0, -1), 3},
	{Copysign(0, -1), Pi},
	{Copysign(0, -1), Inf(1)},

	{0, Inf(-1)},
	{0, -Pi},
	{0, -3},
	{0, Copysign(0, -1)},
	{0, 0},
	{0, 3},
	{0, Pi},
	{0, Inf(1)},
	{0, NaN()},

	{1 / 2, Inf(-1)},
	{1 / 2, Inf(1)},
	{1, Inf(-1)},
	{1, Inf(1)},
	{1, NaN()},

	{Pi, Inf(-1)},
	{Pi, Copysign(0, -1)},
	{Pi, 0},
	{Pi, 1},
	{Pi, Inf(1)},
	{Pi, NaN()},
	{Inf(1), -Pi},
	{Inf(1), Copysign(0, -1)},
	{Inf(1), 0},
	{Inf(1), 1},
	{Inf(1), Pi},
	{Inf(1), NaN()},
	{NaN(), -Pi},
	{NaN(), Copysign(0, -1)},
	{NaN(), 0},
	{NaN(), 1},
	{NaN(), Pi},
	{NaN(), NaN()},
}
var powSC = []float32{
	0,               // pow(-Inf, -Pi)
	Copysign(0, -1), // pow(-Inf, -3)
	1,               // pow(-Inf, -0)
	1,               // pow(-Inf, +0)
	Inf(-1),         // pow(-Inf, 1)
	Inf(-1),         // pow(-Inf, 3)
	Inf(1),          // pow(-Inf, Pi)
	NaN(),           // pow(-Inf, NaN)
	0,               // pow(-Pi, -Inf)
	NaN(),           // pow(-Pi, -Pi)
	1,               // pow(-Pi, -0)
	1,               // pow(-Pi, +0)
	-Pi,             // pow(-Pi, 1)
	NaN(),           // pow(-Pi, Pi)
	Inf(1),          // pow(-Pi, +Inf)
	NaN(),           // pow(-Pi, NaN)
	1,               // pow(-1, -Inf) IEEE 754-2008
	1,               // pow(-1, +Inf) IEEE 754-2008
	NaN(),           // pow(-1, NaN)
	Inf(1),          // pow(-1/2, -Inf)
	0,               // pow(-1/2, +Inf)
	Inf(1),          // pow(-0, -Inf)
	Inf(1),          // pow(-0, -Pi)
	Inf(-1),         // pow(-0, -3) IEEE 754-2008
	Copysign(0, -1), // pow(-0, 3) IEEE 754-2008
	0,               // pow(-0, +Pi)
	0,               // pow(-0, +Inf)
	Inf(1),          // pow(+0, -Inf)
	Inf(1),          // pow(+0, -Pi)
	Inf(1),          // pow(+0, -3)
	1,               // pow(+0, -0)
	1,               // pow(+0, +0)
	0,               // pow(+0, 3)
	0,               // pow(+0, +Pi)
	0,               // pow(+0, +Inf)
	NaN(),           // pow(+0, NaN)
	Inf(1),          // pow(1/2, -Inf)
	0,               // pow(1/2, +Inf)
	1,               // pow(1, -Inf) IEEE 754-2008
	1,               // pow(1, +Inf) IEEE 754-2008
	1,               // pow(1, NaN) IEEE 754-2008
	0,               // pow(+Pi, -Inf)
	1,               // pow(+Pi, -0)
	1,               // pow(+Pi, +0)
	Pi,              // pow(+Pi, 1)
	Inf(1),          // pow(+Pi, +Inf)
	NaN(),           // pow(+Pi, NaN)
	0,               // pow(+Inf, -Pi)
	1,               // pow(+Inf, -0)
	1,               // pow(+Inf, +0)
	Inf(1),          // pow(+Inf, 1)
	Inf(1),          // pow(+Inf, Pi)
	NaN(),           // pow(+Inf, NaN)
	NaN(),           // pow(NaN, -Pi)
	1,               // pow(NaN, -0)
	1,               // pow(NaN, +0)
	NaN(),           // pow(NaN, 1)
	NaN(),           // pow(NaN, +Pi)
	NaN(),           // pow(NaN, NaN)
}

var vfpow10SC = []int{
	math.MinInt32,
	math.MaxInt32,
	-325,
	309,
}

var pow10SC = []float32{
	0,      // pow10(MinInt32)
	Inf(1), // pow10(MaxInt32)
	0,      // pow10(-325)
	Inf(1), // pow10(309)
}

var vfsignbitSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var signbitSC = []bool{
	true,
	true,
	false,
	false,
	false,
}

var vfsinSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var sinSC = []float32{
	NaN(),
	Copysign(0, -1),
	0,
	NaN(),
	NaN(),
}

var vfsinhSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var sinhSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}

var vfsqrtSC = []float32{
	Inf(-1),
	-Pi,
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
	math.Float32frombits(2), // subnormal; see https://golang.org/issue/13013
}
var sqrtSC = []float32{
	NaN(),
	NaN(),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
	5.293956e-23,
}

var vftanhSC = []float32{
	Inf(-1),
	Copysign(0, -1),
	0,
	Inf(1),
	NaN(),
}
var tanhSC = []float32{
	-1,
	Copysign(0, -1),
	0,
	1,
	NaN(),
}

var vfy0SC = []float32{
	Inf(-1),
	0,
	Inf(1),
	NaN(),
}
var y0SC = []float32{
	NaN(),
	Inf(-1),
	0,
	NaN(),
}
var y1SC = []float32{
	NaN(),
	Inf(-1),
	0,
	NaN(),
}
var y2SC = []float32{
	NaN(),
	Inf(-1),
	0,
	NaN(),
}
var yM3SC = []float32{
	NaN(),
	Inf(1),
	0,
	NaN(),
}

// arguments and expected results for boundary cases
//const (
//SmallestNormalFloat64   = 2.2250738585072014e-308 // 2**-1022
//LargestSubnormalFloat64 = SmallestNormalFloat64 - SmallestNonzeroFloat64
//)

var vffrexpBC = []float32{
//SmallestNormalFloat64,
//LargestSubnormalFloat64,
//SmallestNonzeroFloat64,
//MaxFloat64,
//-SmallestNormalFloat64,
//-LargestSubnormalFloat64,
//-SmallestNonzeroFloat64,
//-MaxFloat64,
}
var frexpBC = []fi{
	{0.5, -1021},
	{0.99999999999999978, -1022},
	{0.5, -1073},
	{0.99999999999999989, 1024},
	{-0.5, -1021},
	{-0.99999999999999978, -1022},
	{-0.5, -1073},
	{-0.99999999999999989, 1024},
}

var vfldexpBC = []fi{
	//{SmallestNormalFloat64, -52},
	//{LargestSubnormalFloat64, -51},
	//{SmallestNonzeroFloat64, 1074},
	//{MaxFloat64, -(1023 + 1074)},
	{1, -1075},
	{-1, -1075},
	{1, 1024},
	{-1, 1024},
}
var ldexpBC = []float32{
	//SmallestNonzeroFloat64,
	1e-323, // 2**-1073
	1,
	1e-323, // 2**-1073
	0,
	Copysign(0, -1),
	Inf(1),
	Inf(-1),
}

var logbBC = []float32{
	-1022,
	-1023,
	-1074,
	1023,
	-1022,
	-1023,
	-1074,
	1023,
}

func tolerance(a, b, e float32) bool {
	// Multiplying by e here can underflow denormal values to zero.
	// Check a==b so that at least if a and b are small and identical
	// we say they match.
	if a == b {
		return true
	}
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}
func close(a, b float32) bool      { return tolerance(a, b, 1e-14) }
func veryclose(a, b float32) bool  { return tolerance(a, b, 4e-16) }
func soclose(a, b, e float32) bool { return tolerance(a, b, e) }
func alike(a, b float32) bool {
	switch {
	case IsNaN(a) && IsNaN(b):
		return true
	case a == b:
		return Signbit(a) == Signbit(b)
	}
	return false
}

func TestNaN(t *testing.T) {
	f32 := NaN()
	if f32 == f32 {
		t.Fatalf("NaN() returns %g, expected NaN", f32)
	}
}

func TestAcos(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 10
		if f := Acos(a); !close(acos[i], f) {
			t.Errorf("Acos(%g) = %g, want %g", a, f, acos[i])
		}
	}
	for i := 0; i < len(vfacosSC); i++ {
		if f := Acos(vfacosSC[i]); !alike(acosSC[i], f) {
			t.Errorf("Acos(%g) = %g, want %g", vfacosSC[i], f, acosSC[i])
		}
	}
}

func TestAcosh(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := 1 + Abs(vf[i])
		if f := Acosh(a); !veryclose(acosh[i], f) {
			t.Errorf("Acosh(%g) = %g, want %g", a, f, acosh[i])
		}
	}
	for i := 0; i < len(vfacoshSC); i++ {
		if f := Acosh(vfacoshSC[i]); !alike(acoshSC[i], f) {
			t.Errorf("Acosh(%g) = %g, want %g", vfacoshSC[i], f, acoshSC[i])
		}
	}
}

func TestAsin(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 10
		if f := Asin(a); !veryclose(asin[i], f) {
			t.Errorf("Asin(%g) = %g, want %g", a, f, asin[i])
		}
	}
	for i := 0; i < len(vfasinSC); i++ {
		if f := Asin(vfasinSC[i]); !alike(asinSC[i], f) {
			t.Errorf("Asin(%g) = %g, want %g", vfasinSC[i], f, asinSC[i])
		}
	}
}

func TestAsinh(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Asinh(vf[i]); !veryclose(asinh[i], f) {
			t.Errorf("Asinh(%g) = %g, want %g", vf[i], f, asinh[i])
		}
	}
	for i := 0; i < len(vfasinhSC); i++ {
		if f := Asinh(vfasinhSC[i]); !alike(asinhSC[i], f) {
			t.Errorf("Asinh(%g) = %g, want %g", vfasinhSC[i], f, asinhSC[i])
		}
	}
}

func TestAtan(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Atan(vf[i]); !veryclose(atan[i], f) {
			t.Errorf("Atan(%g) = %g, want %g", vf[i], f, atan[i])
		}
	}
	for i := 0; i < len(vfatanSC); i++ {
		if f := Atan(vfatanSC[i]); !alike(atanSC[i], f) {
			t.Errorf("Atan(%g) = %g, want %g", vfatanSC[i], f, atanSC[i])
		}
	}
}

func TestAtanh(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 10
		if f := Atanh(a); !veryclose(atanh[i], f) {
			t.Errorf("Atanh(%g) = %g, want %g", a, f, atanh[i])
		}
	}
	for i := 0; i < len(vfatanhSC); i++ {
		if f := Atanh(vfatanhSC[i]); !alike(atanhSC[i], f) {
			t.Errorf("Atanh(%g) = %g, want %g", vfatanhSC[i], f, atanhSC[i])
		}
	}
}

func TestAtan2(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Atan2(10, vf[i]); !veryclose(atan2[i], f) {
			t.Errorf("Atan2(10, %g) = %g, want %g", vf[i], f, atan2[i])
		}
	}
	for i := 0; i < len(vfatan2SC); i++ {
		if f := Atan2(vfatan2SC[i][0], vfatan2SC[i][1]); !alike(atan2SC[i], f) {
			t.Errorf("Atan2(%g, %g) = %g, want %g", vfatan2SC[i][0], vfatan2SC[i][1], f, atan2SC[i])
		}
	}
}

func TestCbrt(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Cbrt(vf[i]); !veryclose(cbrt[i], f) {
			t.Errorf("Cbrt(%g) = %g, want %g", vf[i], f, cbrt[i])
		}
	}
	for i := 0; i < len(vfcbrtSC); i++ {
		if f := Cbrt(vfcbrtSC[i]); !alike(cbrtSC[i], f) {
			t.Errorf("Cbrt(%g) = %g, want %g", vfcbrtSC[i], f, cbrtSC[i])
		}
	}
}

func TestCeil(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Ceil(vf[i]); ceil[i] != f {
			t.Errorf("Ceil(%g) = %g, want %g", vf[i], f, ceil[i])
		}
	}
	for i := 0; i < len(vfceilSC); i++ {
		if f := Ceil(vfceilSC[i]); !alike(ceilSC[i], f) {
			t.Errorf("Ceil(%g) = %g, want %g", vfceilSC[i], f, ceilSC[i])
		}
	}
}

func TestCopysign(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Copysign(vf[i], -1); copysign[i] != f {
			t.Errorf("Copysign(%g, -1) = %g, want %g", vf[i], f, copysign[i])
		}
	}
	for i := 0; i < len(vf); i++ {
		if f := Copysign(vf[i], 1); -copysign[i] != f {
			t.Errorf("Copysign(%g, 1) = %g, want %g", vf[i], f, -copysign[i])
		}
	}
	for i := 0; i < len(vfcopysignSC); i++ {
		if f := Copysign(vfcopysignSC[i], -1); !alike(copysignSC[i], f) {
			t.Errorf("Copysign(%g, -1) = %g, want %g", vfcopysignSC[i], f, copysignSC[i])
		}
	}
}

func TestCos(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Cos(vf[i]); !veryclose(cos[i], f) {
			t.Errorf("Cos(%g) = %g, want %g", vf[i], f, cos[i])
		}
	}
	for i := 0; i < len(vfcosSC); i++ {
		if f := Cos(vfcosSC[i]); !alike(cosSC[i], f) {
			t.Errorf("Cos(%g) = %g, want %g", vfcosSC[i], f, cosSC[i])
		}
	}
}

func TestCosh(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Cosh(vf[i]); !close(cosh[i], f) {
			t.Errorf("Cosh(%g) = %g, want %g", vf[i], f, cosh[i])
		}
	}
	for i := 0; i < len(vfcoshSC); i++ {
		if f := Cosh(vfcoshSC[i]); !alike(coshSC[i], f) {
			t.Errorf("Cosh(%g) = %g, want %g", vfcoshSC[i], f, coshSC[i])
		}
	}
}

func TestErf(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 10
		if f := Erf(a); !veryclose(erf[i], f) {
			t.Errorf("Erf(%g) = %g, want %g", a, f, erf[i])
		}
	}
	for i := 0; i < len(vferfSC); i++ {
		if f := Erf(vferfSC[i]); !alike(erfSC[i], f) {
			t.Errorf("Erf(%g) = %g, want %g", vferfSC[i], f, erfSC[i])
		}
	}
}

func TestErfc(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 10
		if f := Erfc(a); !veryclose(erfc[i], f) {
			t.Errorf("Erfc(%g) = %g, want %g", a, f, erfc[i])
		}
	}
	for i := 0; i < len(vferfcSC); i++ {
		if f := Erfc(vferfcSC[i]); !alike(erfcSC[i], f) {
			t.Errorf("Erfc(%g) = %g, want %g", vferfcSC[i], f, erfcSC[i])
		}
	}
}

func TestExp(t *testing.T) {
	testExp(t, Exp, "Exp")
}

func testExp(t *testing.T, Exp func(float32) float32, name string) {
	for i := 0; i < len(vf); i++ {
		if f := Exp(vf[i]); !veryclose(exp[i], f) {
			t.Errorf("%s(%g) = %g, want %g", name, vf[i], f, exp[i])
		}
	}
	for i := 0; i < len(vfexpSC); i++ {
		if f := Exp(vfexpSC[i]); !alike(expSC[i], f) {
			t.Errorf("%s(%g) = %g, want %g", name, vfexpSC[i], f, expSC[i])
		}
	}
}

func TestExpm1(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := vf[i] / 100
		if f := Expm1(a); !veryclose(expm1[i], f) {
			t.Errorf("Expm1(%g) = %g, want %g", a, f, expm1[i])
		}
	}
	for i := 0; i < len(vf); i++ {
		a := vf[i] * 10
		if f := Expm1(a); !close(expm1Large[i], f) {
			t.Errorf("Expm1(%g) = %g, want %g", a, f, expm1Large[i])
		}
	}
	for i := 0; i < len(vfexpm1SC); i++ {
		if f := Expm1(vfexpm1SC[i]); !alike(expm1SC[i], f) {
			t.Errorf("Expm1(%g) = %g, want %g", vfexpm1SC[i], f, expm1SC[i])
		}
	}
}

func TestExp2(t *testing.T) {
	testExp2(t, Exp2, "Exp2")
}

func testExp2(t *testing.T, Exp2 func(float32) float32, name string) {
	for i := 0; i < len(vf); i++ {
		if f := Exp2(vf[i]); !close(exp2[i], f) {
			t.Errorf("%s(%g) = %g, want %g", name, vf[i], f, exp2[i])
		}
	}
	for i := 0; i < len(vfexpSC); i++ {
		if f := Exp2(vfexpSC[i]); !alike(expSC[i], f) {
			t.Errorf("%s(%g) = %g, want %g", name, vfexpSC[i], f, expSC[i])
		}
	}
	for n := -1074; n < 1024; n++ {
		f := Exp2(float32(n))
		vf := Ldexp(1, n)
		if f != vf {
			t.Errorf("%s(%d) = %g, want %g", name, n, f, vf)
		}
	}
}

func TestAbs(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Abs(vf[i]); fabs[i] != f {
			t.Errorf("Abs(%g) = %g, want %g", vf[i], f, fabs[i])
		}
	}
	for i := 0; i < len(vffabsSC); i++ {
		if f := Abs(vffabsSC[i]); !alike(fabsSC[i], f) {
			t.Errorf("Abs(%g) = %g, want %g", vffabsSC[i], f, fabsSC[i])
		}
	}
}

func TestDim(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Dim(vf[i], 0); fdim[i] != f {
			t.Errorf("Dim(%g, %g) = %g, want %g", vf[i], 0.0, f, fdim[i])
		}
	}
	for i := 0; i < len(vffdimSC); i++ {
		if f := Dim(vffdimSC[i][0], vffdimSC[i][1]); !alike(fdimSC[i], f) {
			t.Errorf("Dim(%g, %g) = %g, want %g", vffdimSC[i][0], vffdimSC[i][1], f, fdimSC[i])
		}
	}
	for i := 0; i < len(vffdim2SC); i++ {
		if f := Dim(vffdim2SC[i][0], vffdim2SC[i][1]); !alike(fdimSC[i], f) {
			t.Errorf("Dim(%g, %g) = %g, want %g", vffdim2SC[i][0], vffdim2SC[i][1], f, fdimSC[i])
		}
	}
}

func TestFloor(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Floor(vf[i]); floor[i] != f {
			t.Errorf("Floor(%g) = %g, want %g", vf[i], f, floor[i])
		}
	}
	for i := 0; i < len(vfceilSC); i++ {
		if f := Floor(vfceilSC[i]); !alike(ceilSC[i], f) {
			t.Errorf("Floor(%g) = %g, want %g", vfceilSC[i], f, ceilSC[i])
		}
	}
}

func TestMax(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Max(vf[i], ceil[i]); ceil[i] != f {
			t.Errorf("Max(%g, %g) = %g, want %g", vf[i], ceil[i], f, ceil[i])
		}
	}
	for i := 0; i < len(vffdimSC); i++ {
		if f := Max(vffdimSC[i][0], vffdimSC[i][1]); !alike(fmaxSC[i], f) {
			t.Errorf("Max(%g, %g) = %g, want %g", vffdimSC[i][0], vffdimSC[i][1], f, fmaxSC[i])
		}
	}
	for i := 0; i < len(vffdim2SC); i++ {
		if f := Max(vffdim2SC[i][0], vffdim2SC[i][1]); !alike(fmaxSC[i], f) {
			t.Errorf("Max(%g, %g) = %g, want %g", vffdim2SC[i][0], vffdim2SC[i][1], f, fmaxSC[i])
		}
	}
}

func TestMin(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Min(vf[i], floor[i]); floor[i] != f {
			t.Errorf("Min(%g, %g) = %g, want %g", vf[i], floor[i], f, floor[i])
		}
	}
	for i := 0; i < len(vffdimSC); i++ {
		if f := Min(vffdimSC[i][0], vffdimSC[i][1]); !alike(fminSC[i], f) {
			t.Errorf("Min(%g, %g) = %g, want %g", vffdimSC[i][0], vffdimSC[i][1], f, fminSC[i])
		}
	}
	for i := 0; i < len(vffdim2SC); i++ {
		if f := Min(vffdim2SC[i][0], vffdim2SC[i][1]); !alike(fminSC[i], f) {
			t.Errorf("Min(%g, %g) = %g, want %g", vffdim2SC[i][0], vffdim2SC[i][1], f, fminSC[i])
		}
	}
}

func TestMod(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Mod(10, vf[i]); fmod[i] != f {
			t.Errorf("Mod(10, %g) = %g, want %g", vf[i], f, fmod[i])
		}
	}
	for i := 0; i < len(vffmodSC); i++ {
		if f := Mod(vffmodSC[i][0], vffmodSC[i][1]); !alike(fmodSC[i], f) {
			t.Errorf("Mod(%g, %g) = %g, want %g", vffmodSC[i][0], vffmodSC[i][1], f, fmodSC[i])
		}
	}
}

func TestFrexp(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f, j := Frexp(vf[i]); !veryclose(frexp[i].f, f) || frexp[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vf[i], f, j, frexp[i].f, frexp[i].i)
		}
	}
	for i := 0; i < len(vffrexpSC); i++ {
		if f, j := Frexp(vffrexpSC[i]); !alike(frexpSC[i].f, f) || frexpSC[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vffrexpSC[i], f, j, frexpSC[i].f, frexpSC[i].i)
		}
	}
	for i := 0; i < len(vffrexpBC); i++ {
		if f, j := Frexp(vffrexpBC[i]); !alike(frexpBC[i].f, f) || frexpBC[i].i != j {
			t.Errorf("Frexp(%g) = %g, %d, want %g, %d", vffrexpBC[i], f, j, frexpBC[i].f, frexpBC[i].i)
		}
	}
}

func TestGamma(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Gamma(vf[i]); !close(gamma[i], f) {
			t.Errorf("Gamma(%g) = %g, want %g", vf[i], f, gamma[i])
		}
	}
	for _, g := range vfgamma {
		f := Gamma(g[0])
		var ok bool
		if IsNaN(g[1]) || IsInf(g[1], 0) || g[1] == 0 || f == 0 {
			ok = alike(g[1], f)
		} else if g[0] > -50 && g[0] <= 171 {
			ok = veryclose(g[1], f)
		} else {
			ok = close(g[1], f)
		}
		if !ok {
			t.Errorf("Gamma(%g) = %g, want %g", g[0], f, g[1])
		}
	}
}

func TestHypot(t *testing.T) {
	for i := 0; i < len(vfhypotSC); i++ {
		if f := Hypot(vfhypotSC[i][0], vfhypotSC[i][1]); !alike(hypotSC[i], f) {
			t.Errorf("Hypot(%g, %g) = %g, want %g", vfhypotSC[i][0], vfhypotSC[i][1], f, hypotSC[i])
		}
	}
}

func TestSqrt(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := Abs(vf[i])
		if f := Sqrt(a); sqrt[i] != f {
			t.Errorf("Sqrt(%g) = %g, want %g", a, f, sqrt[i])
		}
	}
	for i := 0; i < len(vfsqrtSC); i++ {
		if f := Sqrt(vfsqrtSC[i]); !alike(sqrtSC[i], f) {
			t.Errorf("Sqrt(%g) = %g, want %g", vfsqrtSC[i], f, sqrtSC[i])
		}
	}
}

func TestLog(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		a := Abs(vf[i])
		if f := Log(a); log[i] != f {
			t.Errorf("Log(%g) = %g, want %g", a, f, log[i])
		}
	}
	if f := Log(10); f != Ln10 {
		t.Errorf("Log(%g) = %g, want %g", 10.0, f, Ln10)
	}
	for i := 0; i < len(vflogSC); i++ {
		if f := Log(vflogSC[i]); !alike(logSC[i], f) {
			t.Errorf("Log(%g) = %g, want %g", vflogSC[i], f, logSC[i])
		}
	}
}

func TestPow(t *testing.T) {
	for i := 0; i < len(vf); i++ {
		if f := Pow(10, vf[i]); !close(pow[i], f) {
			t.Errorf("Pow(10, %g) = %g, want %g", vf[i], f, pow[i])
		}
	}
	for i := 0; i < len(vfpowSC); i++ {
		if f := Pow(vfpowSC[i][0], vfpowSC[i][1]); !alike(powSC[i], f) {
			t.Errorf("Pow(%g, %g) = %g, want %g", vfpowSC[i][0], vfpowSC[i][1], f, powSC[i])
		}
	}
}

func TestPow10(t *testing.T) {
	for i := 0; i < len(vfpow10SC); i++ {
		if f := Pow10(vfpow10SC[i]); !alike(pow10SC[i], f) {
			t.Errorf("Pow10(%d) = %g, want %g", vfpow10SC[i], f, pow10SC[i])
		}
	}
}

func TestApprox(t *testing.T) {
	approxTests := []struct {
		x, y float32
		want bool // true means equal
	}{
		{1.0, 1.0, true},
		{1.0, 1.000001, true},
		{1.0, 1.00001, true},
		{1.0, 1.0001, false},
		{1.0, 1.001, false},
		{1.0, 1.01, false},
		{1.0, 0.999999, true},
		{1.0, 0.99999, true},
		{1.0, 0.9999, false},
		{1.0, 0.999, false},
		{1.0, 0.99, false},
		{0.0, 0.000001, true},
		{0.0, 0.00001, true},
		{0.0, 0.0001, false},
		{0.0, 0.001, false},
		{0.0, 0.01, false},
		{0.0, -0.000001, true},
		{0.0, -0.00001, true},
		{0.0, -0.0001, false},
		{0.0, -0.001, false},
		{0.0, -0.01, false},
		{1e12, 1e12 + 0.000001, true},
		{1e12, 1e12 + 0.00001, true},
		{1e12, 1e12 + 0.0001, true},
		{1e12, 1e12 + 0.001, true},
		{1e12, 1e12 + 0.01, true},
		{1e12, 1e12 - 0.000001, true},
		{1e12, 1e12 - 0.00001, true},
		{1e12, 1e12 - 0.0001, true},
		{1e12, 1e12 - 0.001, true},
		{1e12, 1e12 - 0.01, true},
		{NaN(), 0, false},
		{NaN(), NaN(), false},
		// TODO: missing some +Inf/-Inf test cases
	}

	for _, tt := range approxTests {
		got := Approx(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("Approx(%f, %f) == %t, want %t", tt.x, tt.y, got, tt.want)
		}
	}
}

func TestApproxEpsilon(t *testing.T) {
	approxTests := []struct {
		x, y float32
		epsf float32 // epsilon
		want bool    // true means equal
	}{
		// same epsilon value as in Approx
		{1.0, 1.0, epsilon32 * 100, true},
		{1.0, 1.000001, epsilon32 * 100, true},
		{1.0, 1.00001, epsilon32 * 100, true},
		{1.0, 1.0001, epsilon32 * 100, false},
		{1.0, 1.001, epsilon32 * 100, false},
		{1.0, 1.01, epsilon32 * 100, false},
		{1.0, 1.00001, 0.00001, true},
		{1.0, 1.0001, 0.0001, true},
		{1.0, 1.001, 0.001, true},
		{1.0, 1.01, 0.01, true},
		{1.0, 1.1, 0.1, true},
		{10, 9.999990, 0.000001, true},
		{10, 9.999989, 0.000001, false},
		{10, 9.999988, 0.000001, false},
		{10, 10.000012, 0.000001, false},
		{10, 10.000011, 0.000001, false},
		{10, 10.000010, 0.000001, true},
		{NaN(), 0, 0.1, false},
		{NaN(), NaN(), 0.1, false},
	}

	for _, tt := range approxTests {
		got := ApproxEpsilon(tt.x, tt.y, tt.epsf)
		if got != tt.want {
			t.Errorf("ApproxEpsilon(%f, %f, eps=%f) == %t, want %t", tt.x, tt.y, tt.epsf, got, tt.want)
		}
	}
}
