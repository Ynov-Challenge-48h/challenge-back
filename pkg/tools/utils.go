package tools

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

func AnalyzeImage(imagePath string) (string, error) {
	// Générer un nom de fichier aléatoire
	randomSuffix := rand.Intn(1000000)
	outputFile := filepath.Join("./tmp", fmt.Sprintf("output_%d", randomSuffix))

	// Commande Tesseract
	cmd := exec.Command("Tesseract", imagePath, outputFile, "-l", "fra", "--psm", "6")

	// Exécuter la commande
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'exécution de tesseract : %v", err)
	}

	// Lire le contenu du fichier de sortie
	outputContent, err := os.ReadFile(outputFile + ".txt")
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du fichier de sortie : %v", err)
	}

	// Supprimer les fichiers temporaires
	err = os.Remove(imagePath)
	if err != nil {
		log.Printf("Erreur lors de la suppression de l'image : %v", err)
	}

	err = os.Remove(outputFile + ".txt")
	if err != nil {
		log.Printf("Erreur lors de la suppression du fichier de sortie : %v", err)
	}

	// Retourner le contenu du fichier de sortie
	return string(outputContent), nil
}
