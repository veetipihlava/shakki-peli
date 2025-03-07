import React, { useEffect, useState } from "react";
import { PieceProps } from './types'

// Dynamically import all SVG files using Vite's import.meta.glob
const svgs = import.meta.glob('./assets/*/*.svg');

export function ChessPiece(props: PieceProps) {
  const { id, name, color } = props;

  const [pieceSVG, setPieceSVG] = useState<string | null>(null);

  // Map piece to the corresponding SVG file key
  const getPieceSVG = (pieceName: string): string => {
    switch (pieceName) {
      case "K": return "./assets/light/klt.svg";
      case "k": return "./assets/dark/kdt.svg";
      case "Q": return "./assets/light/qlt.svg";
      case "q": return "./assets/dark/qdt.svg";
      case "B": return "./assets/light/blt.svg";
      case "b": return "./assets/dark/bdt.svg";
      case "N": return "./assets/light/nlt.svg";
      case "n": return "./assets/dark/ndt.svg";
      case "R": return "./assets/light/rlt.svg";
      case "r": return "./assets/dark/rdt.svg";
      case "P": return "./assets/light/plt.svg";
      case "p": return "./assets/dark/pdt.svg";
      default: return "";
    }
  };

  useEffect(() => {
    const svgPath = getPieceSVG(name); // Get the SVG path string
    if (svgPath) {
      // Dynamically import the SVG file from the glob
      const importSVG = svgs[svgPath];

      if (importSVG) {
        importSVG().then((module) => {
          const svgUrl = module.default; // Assuming `default` contains the URL
          setPieceSVG(svgUrl);
        }).catch((error) => {
          console.error("Error loading SVG:", error);
          setPieceSVG(null); // Fallback if there's an error
        });
      } else {
        console.error(`SVG path not found for piece "${name}"`);
        setPieceSVG(null); // Fallback if the SVG path doesn't exist
      }
    }
  }, [name]); // Re-run when piece changes

  if (!pieceSVG) {
    return null; // Return null if no SVG is found
  }

  return <img className="chessPiece" src={pieceSVG} alt={name} />;
};
