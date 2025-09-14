package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// CraftType represents the valid craft types.
// https://wiki.openstreetmap.org/wiki/Key:craft
type CraftType string

const (
	// CraftAgriculturalEngines represents a place where engines and tools for agricultural use are planned and constructed.
	CraftAgriculturalEngines CraftType = "agricultural_engines"
	// CraftAtelier represents a place where visual artists work.
	CraftAtelier CraftType = "atelier"
	// CraftBagRepair represents a place where bags/luggage are repaired.
	CraftBagRepair CraftType = "bag_repair"
	// CraftBakery represents a workplace for fresh bakery goods. May have no shop or be combined with a shop.
	CraftBakery CraftType = "bakery"
	// CraftBasketMaker represents a person weaving baskets.
	CraftBasketMaker CraftType = "basket_maker"
	// CraftBeekeeper represents the workplace of a beekeeper (apiarist).
	CraftBeekeeper CraftType = "beekeeper"
	// CraftBlacksmith represents a place where a blacksmith forges tools, horseshoes, etc. from iron.
	CraftBlacksmith CraftType = "blacksmith"
	// CraftBoatbuilder represents a workplace where boats are planned and constructed.
	CraftBoatbuilder CraftType = "boatbuilder"
	// CraftBookbinder represents a workplace for physically assembling a book from a number of sheets of paper.
	CraftBookbinder CraftType = "bookbinder"
	// CraftBrewery represents a dedicated building for the making of beer.
	CraftBrewery CraftType = "brewery"
	// CraftBuilder represents a workplace or office of a tradesman who performs structural alterations and additions to buildings.
	CraftBuilder CraftType = "builder"
	// CraftCabinetMaker represents a person who makes fine wooden furniture.
	CraftCabinetMaker CraftType = "cabinet_maker"
	// CraftCandlemaker represents a person who makes candles or a manufactory where candles were made.
	CraftCandlemaker CraftType = "candlemaker"
	// CraftCarPainter represents a place specializing in painting cars.
	CraftCarPainter CraftType = "car_painter"
	// CraftCarpenter represents a workplace or office of carpenters that work with timber to construct, install and maintain buildings, furniture, and other objects.
	CraftCarpenter CraftType = "carpenter"
	// CraftCarpetCleaner represents a carpet cleaner is a tradesman who specializes in cleaning carpets.
	CraftCarpetCleaner CraftType = "carpet_cleaner"
	// CraftCarpetLayer represents a carpet layer is a tradesmen who specializes in laying carpets.
	CraftCarpetLayer CraftType = "carpet_layer"
	// CraftCaterer represents a workplace or office of one who prepares customized meals for takeout, or provides prepared meals or supplies to a group at social gatherings
	CraftCaterer CraftType = "caterer"
	// CraftChimneySweeper represents a workplace or office of a person who cleans chimneys for a living.
	CraftChimneySweeper CraftType = "chimney_sweeper"
	// CraftCleaning represents workplace or office of a person who cleans buildings, industrial facilities or windows.
	CraftCleaning CraftType = "cleaning"
	// CraftClockmaker represents a workplace or office of a clockmaker that is building, repairing or engraving clocks of all sorts.
	CraftClockmaker CraftType = "clockmaker"
	// CraftClothesMending represents a place where clothing is mended (repaired) or altered professionally.
	CraftClothesMending CraftType = "clothes_mending"
	// CraftConfectionery represents a place where the set of food items that are rich in sugar, any one or type of which is called a confection is produced.
	CraftConfectionery CraftType = "confectionery"
	// CraftCooper represents a person or company that manufacture of containers and vessels mainly made of wood.
	CraftCooper CraftType = "cooper"
	// CraftDentalTechnician represents a workplace for dental technician who constructs custom-made restorative and dental appliances.
	CraftDentalTechnician CraftType = "dental_technician"
	// CraftDistillery represents an establishment for distilling, especially for distilling alcoholic liquors.
	CraftDistillery CraftType = "distillery"
	// CraftDoorConstruction represents a workplace of someone constructing doors.
	CraftDoorConstruction CraftType = "door_construction"
	// CraftDressmaker represents dressmaker is a person who makes custom clothing for women, such as dresses, blouses, and evening gowns.
	CraftDressmaker CraftType = "dressmaker"
	// CraftElectrician represents a person or company that works with electricity or electrical systems.
	CraftElectrician CraftType = "electrician"
	// CraftElectronicsRepair represents repair shops and service centres which provides repair of computers, phones, appliance etc.
	CraftElectronicsRepair CraftType = "electronics_repair"
	// CraftElevator represents company specialized in installing elevators.
	CraftElevator CraftType = "elevator"
	// CraftEmbroiderer represents needlework or textile artist.
	CraftEmbroiderer CraftType = "embroiderer"
	// CraftEngraver represents a workplace or office of an engraver.
	CraftEngraver CraftType = "engraver"
	// CraftFenceMaker represents a workplace or office of a person or company specialized in manufacturing and installing fences.
	CraftFenceMaker CraftType = "fence_maker"
	// CraftFloorer represents a workplace or office of a floorer.
	CraftFloorer CraftType = "floorer"
	// CraftGardener represents a workplace or office of a garden designer or a landscape gardener.
	CraftGardener CraftType = "gardener"
	// CraftGlassblower represents a person or company that blows bottles or other objects from molten glass.
	CraftGlassblower CraftType = "glassblower"
	// CraftGlaziery represents a place where residential, commercial, and artistic glass is selected, cut, installed, replaced, and removed.
	CraftGlaziery CraftType = "glaziery"
	// CraftGoldsmith represents the workplace of a person who forges things out of gold, especially jewelry.
	CraftGoldsmith CraftType = "goldsmith"
	// CraftGrindingMill represents a device that breaks solid materials into smaller pieces by grinding, crushing, or cutting. The aftermath is powdered product, not liquid extract like oil mill.
	CraftGrindingMill CraftType = "grinding_mill"
	// CraftHandicraft represents a place where useful and decorative devices are made completely by hand or using only simple tools.
	CraftHandicraft CraftType = "handicraft"
	// CraftHvac represents a workplace or office of a person or company that install and maintain HVAC systems (Heating, Ventilating, and Air Conditioning).
	CraftHvac CraftType = "hvac"
	// CraftInsulation represents a workplace or office of a person who does thermal insulation in buildings.
	CraftInsulation CraftType = "insulation"
	// CraftInteriorDecorator represents workplace or office of a person who does interior decoratoring.
	CraftInteriorDecorator CraftType = "interior_decorator"
	// CraftInteriorWork represents workplace of a person who builds/installs non-load-bearing interior walls, especially drywall.
	CraftInteriorWork CraftType = "interior_work"
	// CraftJeweller represents a place where necklaces, rings, brooches, earrings and bracelets and other personal adornments are created or repaired.
	CraftJeweller CraftType = "jeweller"
	// CraftJoiner represents an artisan who builds things by joining pieces of wood, particularly furniture or ornamental work.
	CraftJoiner CraftType = "joiner"
	// CraftKeyCutter represents a place where keys can be duplicated from originals.
	CraftKeyCutter CraftType = "key_cutter"
	// CraftLaboratory represents a laboratory.
	CraftLaboratory CraftType = "laboratory"
	// CraftLapidary represents lapidary: a gemcutter.
	CraftLapidary CraftType = "lapidary"
	// CraftLeather represents a leatherworker; a person who makes things from leather.
	CraftLeather CraftType = "leather"
	// CraftLocksmith represents the workshop of a craftsman who installs or repairs locks and opens locked doors.
	CraftLocksmith CraftType = "locksmith"
	// CraftLuthier represents workshop of a luthier, who makes or repairs stringed wooden musical instruments, such as lutes, violins, and guitars.
	CraftLuthier CraftType = "luthier"
	// CraftMetalConstruction represents workplace or office of a person who work with metal. This applies to: planning, construction, trade, repair.
	CraftMetalConstruction CraftType = "metal_construction"
	// CraftMint represents a place where coins and medals are hand crafted.
	CraftMint CraftType = "mint"
	// CraftMusicalInstrument represents a craftsman creating musical instruments
	CraftMusicalInstrument CraftType = "musical_instrument"
	// CraftOilMill represents a mill designed to crush or bruise oil-bearing seeds, such as linseed or peanuts, or other oil-rich vegetable material.
	CraftOilMill CraftType = "oil_mill"
	// CraftOptician represents a place where lenses for the correction of a person's vision are designed, fitted and dispensed.
	CraftOptician CraftType = "optician"
	// CraftOrganBuilder represents organ building is the profession of designing, building, restoring and maintaining pipe organs.
	CraftOrganBuilder CraftType = "organ_builder"
	// CraftPainter represents a workplace or office of a house painter which is a tradesman responsible for the painting and decorating of buildings.
	CraftPainter CraftType = "painter"
	// CraftPaperhanger represents the office of a paperhanger.
	CraftPaperhanger CraftType = "paperhanger"
	// CraftParquetLayer represents a workplace or office of a parquet layer.
	CraftParquetLayer CraftType = "parquet_layer"
	// CraftPaver represents the workshop or office of a paver.
	CraftPaver CraftType = "paver"
	// CraftPestControl represents a business that offers pest control services.
	CraftPestControl CraftType = "pest_control"
	// CraftPhotographer represents a workplace or office of a person who takes photographs using a camera.
	CraftPhotographer CraftType = "photographer"
	// CraftPhotographicLaboratory represents a place where photos taken on a film or digital images are transformed into a more permanent form, often on paper or plastic foil.
	CraftPhotographicLaboratory CraftType = "photographic_laboratory"
	// CraftPhotovoltaic represents a workplace or office who installed or planned photovoltaic systems.
	CraftPhotovoltaic CraftType = "photovoltaic"
	// CraftPianoTuner represents a piano tuner or technician. A professional that tunes, regulates and makes the maintenance of pianos. Most of them repair simple problems and some restore a whole piano.
	CraftPianoTuner CraftType = "piano_tuner"
	// CraftPlasterer represents a workplace or office of a tradesman who works with plaster, such as forming a layer of plaster on an interior wall or plaster decorative moldings on ceilings or walls.
	CraftPlasterer CraftType = "plasterer"
	// CraftPlumber represents a workplace or office of a skilled tradesperson who specializes in (drinking) water supply, sewage and drainage systems.
	CraftPlumber CraftType = "plumber"
	// CraftPottery represents a place where earthenware, stoneware and porcelain and other ceramic ware is made by potters.
	CraftPottery CraftType = "pottery"
	// CraftPrinter represents a small printing business which produces published works such as newspapers, books, magazines, etc.
	CraftPrinter CraftType = "printer"
	// CraftPrintmaker represents a producer of fine art prints.
	CraftPrintmaker CraftType = "printmaker"
	// CraftRestoration represents restorer.
	CraftRestoration CraftType = "restoration"
	// CraftRigger represents the workplace of a person who sets up the rigging for sailboats and sailing ships (or possibly a person who works on ropes, booms, lifts, hoists and the like for a stage production).
	CraftRigger CraftType = "rigger"
	// CraftRoofer represents a workplace or office of a tradesman who is specialized in roof covering and sealing.
	CraftRoofer CraftType = "roofer"
	// CraftSaddler represents a place where saddles or accessories for cars and boats are produced or repaired.
	CraftSaddler CraftType = "saddler"
	// CraftSailmaker represents a workshop of a sailmaker who cleans, repairs and makes sails.
	CraftSailmaker CraftType = "sailmaker"
	// CraftSawmill represents a sawmill is a facility where logs are cut into timber.
	CraftSawmill CraftType = "sawmill"
	// CraftScaffolder represents a workplace or office of a tradesman who builds temporary structure used to support people and material in the construction or repair of buildings and other large structures.
	CraftScaffolder CraftType = "scaffolder"
	// CraftSculptor represents the workplace of a sculptor.
	CraftSculptor CraftType = "sculptor"
	// CraftShoemaker represents a place where shoes, boots, sandals, clogs and moccasins are created, repaired or altered to fit personal needs.
	CraftShoemaker CraftType = "shoemaker"
	// CraftSignmaker represents a maker of signs, mainly for shops and other businesses.
	CraftSignmaker CraftType = "signmaker"
	// CraftStandBuilder represents a workplace of a company that builds stands/booths for fairs.
	CraftStandBuilder CraftType = "stand_builder"
	// CraftStonemason represents a place where rough pieces of rock are shaped into accurate geometrical shapes.
	CraftStonemason CraftType = "stonemason"
	// CraftStoveFitter represents a workplace/office of a person / company who designs and builds stoves / tiled stoves.
	CraftStoveFitter CraftType = "stove_fitter"
	// CraftSunProtection represents a place where sun protection is made, repaired, or delivered for assembly, especially awnings and shutters.
	CraftSunProtection CraftType = "sun_protection"
	// CraftTailor represents a place where clothing is made, repaired, or altered professionally, especially suits and men's clothing.
	CraftTailor CraftType = "tailor"
	// CraftTatami represents a workplace or office of an artisan for tatami mat.
	CraftTatami CraftType = "tatami"
	// CraftTiler represents a workplace or office of a person who lays tiles on floors, swimming pools and such.
	CraftTiler CraftType = "tiler"
	// CraftTinsmith represents a person who makes and repairs things made of light-coloured metal, particularly tinware.
	CraftTinsmith CraftType = "tinsmith"
	// CraftToolmaker represents a workplace of a company that builds complex tools, machine tools, moulds, gauges,... .
	CraftToolmaker CraftType = "toolmaker"
	// CraftTurner represents a workplace or office of a turner.
	CraftTurner CraftType = "turner"
	// CraftUpholsterer represents a place where furniture, especially seats are provided with padding, springs, webbing, and fabric or leather covers.
	CraftUpholsterer CraftType = "upholsterer"
	// CraftWatchmaker represents a watchmaker is an artisan who makes and repairs watches. Today due to industrial production they mostly repair watches.
	CraftWatchmaker CraftType = "watchmaker"
	// CraftWaterWellDrilling represents a craftsman or company who drill water wells.
	CraftWaterWellDrilling CraftType = "water_well_drilling"
	// CraftWeaver represents workshop of a person producing something by weaving, may offer products for sale.
	CraftWeaver CraftType = "weaver"
	// CraftWelder represents one who unites pieces of metal with high temperature.
	CraftWelder CraftType = "welder"
	// CraftWindowConstruction represents a workplace of someone constructing windows.
	CraftWindowConstruction CraftType = "window_construction"
	// CraftWinery represents a place where wine is produced.
	CraftWinery CraftType = "winery"
)

var validCraftTypes = map[string]CraftType{
	"agricultural_engines":    CraftAgriculturalEngines,
	"atelier":                 CraftAtelier,
	"bag_repair":              CraftBagRepair,
	"bakery":                  CraftBakery,
	"basket_maker":            CraftBasketMaker,
	"beekeeper":               CraftBeekeeper,
	"blacksmith":              CraftBlacksmith,
	"boatbuilder":             CraftBoatbuilder,
	"bookbinder":              CraftBookbinder,
	"brewery":                 CraftBrewery,
	"builder":                 CraftBuilder,
	"cabinet_maker":           CraftCabinetMaker,
	"candlemaker":             CraftCandlemaker,
	"car_painter":             CraftCarPainter,
	"carpenter":               CraftCarpenter,
	"carpet_cleaner":          CraftCarpetCleaner,
	"carpet_layer":            CraftCarpetLayer,
	"caterer":                 CraftCaterer,
	"chimney_sweeper":         CraftChimneySweeper,
	"cleaning":                CraftCleaning,
	"clockmaker":              CraftClockmaker,
	"clothes_mending":         CraftClothesMending,
	"confectionery":           CraftConfectionery,
	"cooper":                  CraftCooper,
	"dental_technician":       CraftDentalTechnician,
	"distillery":              CraftDistillery,
	"door_construction":       CraftDoorConstruction,
	"dressmaker":              CraftDressmaker,
	"electrician":             CraftElectrician,
	"electronics_repair":      CraftElectronicsRepair,
	"elevator":                CraftElevator,
	"embroiderer":             CraftEmbroiderer,
	"engraver":                CraftEngraver,
	"fence_maker":             CraftFenceMaker,
	"floorer":                 CraftFloorer,
	"gardener":                CraftGardener,
	"glassblower":             CraftGlassblower,
	"glaziery":                CraftGlaziery,
	"goldsmith":               CraftGoldsmith,
	"grinding_mill":           CraftGrindingMill,
	"handicraft":              CraftHandicraft,
	"hvac":                    CraftHvac,
	"insulation":              CraftInsulation,
	"interior_decorator":      CraftInteriorDecorator,
	"interior_work":           CraftInteriorWork,
	"jeweller":                CraftJeweller,
	"joiner":                  CraftJoiner,
	"key_cutter":              CraftKeyCutter,
	"laboratory":              CraftLaboratory,
	"lapidary":                CraftLapidary,
	"leather":                 CraftLeather,
	"locksmith":               CraftLocksmith,
	"luthier":                 CraftLuthier,
	"metal_construction":      CraftMetalConstruction,
	"mint":                    CraftMint,
	"musical_instrument":      CraftMusicalInstrument,
	"oil_mill":                CraftOilMill,
	"optician":                CraftOptician,
	"organ_builder":           CraftOrganBuilder,
	"painter":                 CraftPainter,
	"paperhanger":             CraftPaperhanger,
	"parquet_layer":           CraftParquetLayer,
	"paver":                   CraftPaver,
	"pest_control":            CraftPestControl,
	"photographer":            CraftPhotographer,
	"photographic_laboratory": CraftPhotographicLaboratory,
	"photovoltaic":            CraftPhotovoltaic,
	"piano_tuner":             CraftPianoTuner,
	"plasterer":               CraftPlasterer,
	"plumber":                 CraftPlumber,
	"pottery":                 CraftPottery,
	"printer":                 CraftPrinter,
	"printmaker":              CraftPrintmaker,
	"restoration":             CraftRestoration,
	"rigger":                  CraftRigger,
	"roofer":                  CraftRoofer,
	"saddler":                 CraftSaddler,
	"sailmaker":               CraftSailmaker,
	"sawmill":                 CraftSawmill,
	"scaffolder":              CraftScaffolder,
	"sculptor":                CraftSculptor,
	"shoemaker":               CraftShoemaker,
	"signmaker":               CraftSignmaker,
	"stand_builder":           CraftStandBuilder,
	"stonemason":              CraftStonemason,
	"stove_fitter":            CraftStoveFitter,
	"sun_protection":          CraftSunProtection,
	"tailor":                  CraftTailor,
	"tatami":                  CraftTatami,
	"tiler":                   CraftTiler,
	"tinsmith":                CraftTinsmith,
	"toolmaker":               CraftToolmaker,
	"turner":                  CraftTurner,
	"upholsterer":             CraftUpholsterer,
	"watchmaker":              CraftWatchmaker,
	"water_well_drilling":     CraftWaterWellDrilling,
	"weaver":                  CraftWeaver,
	"welder":                  CraftWelder,
	"window_construction":     CraftWindowConstruction,
	"winery":                  CraftWinery,
}

// Craft is used to tag a place that produces or processes customised goods.
// https://wiki.openstreetmap.org/wiki/Key:craft
type Craft Feature

func (c *Craft) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validCraftTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validCraftTypes))
			for k := range validCraftTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", c.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = c.GetTag()

	*c = Craft(f)
	return nil
}

func (c *Craft) GetTag() string {
	return "craft"
}

func (c *Craft) ToOQL() string {
	return Feature(*c).ToOQL()
}
