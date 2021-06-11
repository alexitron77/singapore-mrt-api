package models

import (
	"reflect"
	"testing"
)

func TestStationMapping(t *testing.T) {
	expected := map[string]string{"CC1": "Dhoby Ghaut", "CC10": "MacPherson", "CC11": "Tai Seng", "CC12": "Bartley", "CC13": "Serangoon", "CC14": "Lorong Chuan", "CC15": "Bishan", "CC16": "Marymount", "CC17": "Caldecott", "CC19": "Botanic Gardens", "CC2": "Bras Basah", "CC20": "Farrer Road", "CC21": "Holland Village", "CC22": "Buona Vista", "CC23": "one-north", "CC24": "Kent Ridge", "CC25": "Haw Par Villa", "CC26": "Pasir Panjang", "CC27": "Labrador Park", "CC28": "Telok Blangah", "CC29": "HarbourFront", "CC3": "Esplanade", "CC4": "Promenade", "CC5": "Nicoll Highway", "CC6": "Stadium", "CC7": "Mountbatten", "CC8": "Dakota", "CC9": "Paya Lebar", "CE0": "Promenade", "CE1": "Bayfront", "CE2": "Marina Bay", "CG0": "Tanah Merah", "CG1": "Expo", "CG2": "Changi Airport", "DT1": "Bukit Panjang", "DT10": "Stevens", "DT11": "Newton", "DT12": "Little India", "DT13": "Rochor", "DT14": "Bugis", "DT15": "Promenade", "DT16": "Bayfront", "DT17": "Downtown", "DT18": "Telok Ayer", "DT19": "Chinatown", "DT2": "Cashew", "DT20": "Fort Canning", "DT21": "Bencoolen", "DT22": "Jalan Besar", "DT23": "Bendemeer", "DT24": "Geylang Bahru", "DT25": "Mattar", "DT26": "MacPherson", "DT27": "Ubi", "DT28": "Kaki Bukit", "DT29": "Bedok North", "DT3": "Hillview", "DT30": "Bedok Reservoir", "DT31": "Tampines West", "DT32": "Tampines", "DT33": "Tampines East", "DT34": "Upper Changi", "DT35": "Expo", "DT5": "Beauty World", "DT6": "King Albert Park", "DT7": "Sixth Avenue", "DT8": "Tan Kah Kee", "DT9": "Botanic Gardens", "EW1": "Pasir Ris", "EW10": "Kallang", "EW11": "Lavender", "EW12": "Bugis", "EW13": "City Hall", "EW14": "Raffles Place", "EW15": "Tanjong Pagar", "EW16": "Outram Park", "EW17": "Tiong Bahru", "EW18": "Redhill", "EW19": "Queenstown", "EW2": "Tampines", "EW20": "Commonwealth", "EW21": "Buona Vista", "EW22": "Dover", "EW23": "Clementi", "EW24": "Jurong East", "EW25": "Chinese Garden", "EW26": "Lakeside", "EW27": "Boon Lay", "EW28": "Pioneer", "EW29": "Joo Koon", "EW3": "Simei", "EW30": "Gul Circle", "EW31": "Tuas Crescent", "EW32": "Tuas West Road", "EW33": "Tuas Link", "EW4": "Tanah Merah", "EW5": "Bedok", "EW6": "Kembangan", "EW7": "Eunos", "EW8": "Paya Lebar", "EW9": "Aljunied", "NE1": "HarbourFront", "NE10": "Potong Pasir", "NE11": "Woodleigh", "NE12": "Serangoon", "NE13": "Kovan", "NE14": "Hougang", "NE15": "Buangkok", "NE16": "Sengkang", "NE17": "Punggol", "NE3": "Outram Park", "NE4": "Chinatown", "NE5": "Clarke Quay", "NE6": "Dhoby Ghaut", "NE7": "Little India", "NE8": "Farrer Park", "NE9": "Boon Keng", "NS1": "Jurong East", "NS10": "Admiralty", "NS11": "Sembawang", "NS12": "Canberra", "NS13": "Yishun", "NS14": "Khatib", "NS15": "Yio Chu Kang", "NS16": "Ang Mo Kio", "NS17": "Bishan", "NS18": "Braddell", "NS19": "Toa Payoh", "NS2": "Bukit Batok", "NS20": "Novena", "NS21": "Newton", "NS22": "Orchard", "NS23": "Somerset", "NS24": "Dhoby Ghaut", "NS25": "City Hall", "NS26": "Raffles Place", "NS27": "Marina Bay", "NS28": "Marina South Pier", "NS3": "Bukit Gombak", "NS4": "Choa Chu Kang", "NS5": "Yew Tee", "NS7": "Kranji", "NS8": "Marsiling", "NS9": "Woodlands", "TE1": "Woodlands North", "TE10": "Mount Pleasant", "TE11": "Stevens", "TE12": "Napier", "TE13": "Orchard Boulevard", "TE14": "Orchard", "TE15": "Great World", "TE16": "Havelock", "TE17": "Outram Park", "TE18": "Maxwell", "TE19": "Shenton Way", "TE2": "Woodlands", "TE20": "Marina Bay", "TE21": "Marina South", "TE22": "Gardens by the Bay", "TE3": "Woodlands South", "TE4": "Springleaf", "TE5": "Lentor", "TE6": "Mayflower", "TE7": "Bright Hill", "TE8": "Upper Thomson", "TE9": "Caldecott"}

	path := "../../public/station_map.csv"

	res := StnMapping(path)

	if reflect.DeepEqual(expected, res) == false {
		t.Errorf("Expected %v\n Got %v", expected, res)
	} else {
		t.Logf("Successfully pass")
	}
}
