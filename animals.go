package main

// Animal stores information about an animal.
type Animal struct {
	Name  string
	ScientificName string
	AnimalGroup string
	Habitat string
	Diet string
	Lifespan string
	Size string
	ConservationStatus string
	FunFact string
    Image string
}

var animals = []Animal{

	// Tiger
{
    Name: "Tiger",
    ScientificName: "Panthera tigris",
    AnimalGroup: "Mammal",
    Habitat: "Forests, grasslands and mangroves",
    Diet: "Mainly deer, wild boar and other medium-sized animals",
    Lifespan: "10–15 years in the wild",
    Size: "Up to about 3.3 metres long",
    ConservationStatus: "Endangered",
    FunFact: "Every tiger has a unique stripe pattern.",
    Image: "/static/Images/Tiger.jpg",
},

// Blue Whale
{
    Name: "Blue Whale",
    ScientificName: "Balaenoptera musculus",
    AnimalGroup: "Mammal",
    Habitat: "Oceans around the world",
    Diet: "Mainly krill and other tiny crustaceans",
    Lifespan: "About 70–90 years",
    Size: "Up to about 30 metres long",
    ConservationStatus: "Endangered",
    FunFact: "The blue whale is the largest known animal to have ever lived.",
    Image: "/static/Images/Whale.jpg",
},

// Honey Bee
{
    Name: "Honey Bee",
    ScientificName: "Apis mellifera",
    AnimalGroup: "Insect",
    Habitat: "Grasslands, forests, gardens and farms",
    Diet: "Nectar and pollen from flowers",
    Lifespan: "Several weeks to months for worker bees",
    Size: "About 1–1.5 centimetres long",
    ConservationStatus: "Not Evaluated",
    FunFact: "Honey bees can communicate the location of food using a waggle dance.",
    Image: "/static/Images/Bee.jpg",
},

// Bald Eagle
{
    Name: "Bald Eagle",
    ScientificName: "Haliaeetus leucocephalus",
    AnimalGroup: "Bird",
    Habitat: "Forests near lakes, rivers and coasts",
    Diet: "Mainly fish, but also birds, mammals and carrion",
    Lifespan: "About 20–30 years in the wild",
    Size: "About 70–100 centimetres long",
    ConservationStatus: "Least Concern",
    FunFact: "Bald eagles have powerful talons for catching prey.",
    Image: "/static/Images/Eagle.jpg",
},

// Red-Eyed Tree Frog
{
    Name: "Red-Eyed Tree Frog",
    ScientificName: "Agalychnis callidryas",
    AnimalGroup: "Amphibian",
    Habitat: "Tropical rainforests near freshwater",
    Diet: "Mainly insects and other small invertebrates",
    Lifespan: "About 5 years in the wild",
    Size: "About 4–7 centimetres long",
    ConservationStatus: "Least Concern",
    FunFact: "Its bright red eyes can startle predators.",
    Image: "/static/Images/Frog.png",
},

// Gopher
{
    Name: "Gopher",
    ScientificName: "Geomys bursarius",
    AnimalGroup: "Mammal",
    Habitat: "Grasslands and open areas with suitable soil",
    Diet: "Roots, tubers, grasses and other plant material",
    Lifespan: "Up to about 7 years",
    Size: "About 19–36 centimetres long",
    ConservationStatus: "Least Concern",
    FunFact: "Gophers spend much of their lives underground in burrow systems.",
    Image: "/static/Images/Gopher.jpg",
},

// Great White Shark
{
    Name: "Great White Shark",
    ScientificName: "Carcharodon carcharias",
    AnimalGroup: "Fish",
    Habitat: "Coastal and offshore ocean waters",
    Diet: "Fish, rays, seals, sea lions, turtles and other marine animals",
    Lifespan: "About 60–70 years",
    Size: "Up to around 6 metres long",
    ConservationStatus: "Vulnerable",
    FunFact: "Great white sharks can detect tiny amounts of chemicals in water.",
    Image: "/static/Images/Shark.jpg",
},

// Saltwater Crocodile
{
    Name: "Crocodile",
    ScientificName: "Crocodylus porosus",
    AnimalGroup: "Reptile",
    Habitat: "Rivers, wetlands, estuaries and coastal waters",
    Diet: "Fish, birds, mammals, reptiles and other animals",
    Lifespan: "About 70 years or more",
    Size: "Up to about 6 metres long",
    ConservationStatus: "Least Concern",
    FunFact: "Saltwater crocodiles are the largest living reptiles.",
    Image: "/static/Images/Croc.jpg",
},
}
