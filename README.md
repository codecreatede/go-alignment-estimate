go-whole-genome-estimates

- a whole genome linear block estimation. 
- block is defined as a single line for the aligned genome. 
- no matter how many genome you aligned you can pass the alignment file.
- the first genome will be taken as a reference. 
- check the AlignmentGO package for defining your own block size. 

```
[gauravsablok@fedora]~/Desktop/codecreatede/whole-genome-alignment-estimate% \
go run main.go -alignment ./samplefile/samplealignment.fasta
The collinearity block for A as a base pattern and T as a mismatch is 1
The collinearity block for A as a base pattern G as a mismatch is  1
The collinearity block for A as a base pattern and C as a mismatch is 1
The collinearity block for T as a base pattern and G as a mismatch is 0
The collinearity block for T as a base pattern C as a mismatch is  0
The collinearity block for T as a base pattern and A as a mismatch is 0
The collinearity block for G as a base pattern and C as a mismatch is 1
The collinearity block for G as a base pattern A as a mismatch is  0
The collinearity block for G as a base pattern and T as a mismatch is 0
The collinearity block for C as a base pattern and A as a mismatch is 0
The collinearity block for C as a base pattern T as a mismatch is  0
The collinearity block for C as a base pattern and G as a mismatch is 0
```

Gaurav Sablok
