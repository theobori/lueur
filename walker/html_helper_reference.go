package walker

func (w *Walker) walkHTMLReferenceHelper(alt string, destination string) (string, error) {
	if alt == "" {
		alt = destination
	}

	line, err := w.referenceLine(alt, destination)
	if err != nil {
		return "", err
	}

	inlineText := w.processReferenceLineEdgeCases(line, destination)

	w.ctx.ReferencesQueue = append(w.ctx.ReferencesQueue, *line)

	return inlineText, nil
}
