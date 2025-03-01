// Helper function to flatten and collect all odds for a given prop type
function getAllOdds(player, propType, bookMaker) {
    if (!player.odds || !player.odds[propType]) {
        return [];
    }

    if (bookMaker === "All") {
        return Object.values(player.odds[propType]).flat();
    } else if (player.odds[propType][bookMaker]) {
        return player.odds[propType][bookMaker];
    }

    return [];
}

// Helper function to get the most frequent value in an array
function getMostFrequentValue(arr) {
    if (arr.length === 0) return '-';

    const counts = arr.reduce((acc, val) => {
        acc[val] = (acc[val] || 0) + 1;
        return acc;
    }, {});

    return Object.keys(counts)
        .map(Number)
        .sort((a, b) => counts[b] - counts[a])[0] || '-';
}

// Get the most common point value across all bookmakers
export function getMostCommonPoint(player, propType) {
    const allOdds = getAllOdds(player, propType, "All").map(odd => odd.point);
    return getMostFrequentValue(allOdds);
}

// Helper function to get the best price for a given type (Over/Under)
function getBestPrice(player, propType, bookMaker, betType) {
    const allOdds = getAllOdds(player, propType, bookMaker);
    if (allOdds.length === 0) return '-';

    return allOdds
        .filter(odd => odd.name === betType)
        .map(odd => odd.price)
        .sort((a, b) => b - a)[0] || '-';
}

// Get the best "Over" price from all odds
export function getBestOver(player, propType, bookMaker) {
    return getBestPrice(player, propType, bookMaker, 'Over');
}

// Get the best "Under" price from all odds
export function getBestUnder(player, propType, bookMaker) {
    return getBestPrice(player, propType, bookMaker, 'Under');
}
